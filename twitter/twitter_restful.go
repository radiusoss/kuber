package twitter

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/datalayer/kuber/config"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

type TwitterUser struct {
	Username       string `json:"username"`
	Oauth_token    string `json:"oauth_token"`
	Oauth_verifier string `json:"oauth_verifier"`
	IsAuth         bool   `json:"isAuth"`
}

type TwitterResource struct {
}

var SessionName = "kuber-twitter"
var SessionKey = "twitter-state"

var Store = sessions.NewCookieStore([]byte("twitter-session-secret"))

var cff config.Config

var TwitterClient *ServerClient

func (t TwitterResource) WebService(cf config.Config) *restful.WebService {

	cff = cf

	gob.Register(TwitterUser{})

	err := viper.Unmarshal(&config.KuberConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	log.Println("Kuber Config:", config.KuberConfig)

	TwitterClient = NewServerClient(cf.TwitterConsumerKey, cf.TwitterConsumerSecret)

	Store.Options = &sessions.Options{
		Path:     "/",      // to match all requests
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: false,
	}

	ws := new(restful.WebService)
	ws.Path("/api/v1/twitter").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(t.RedirectUserToTwitterRestful))
	ws.Route(ws.GET("/callback").To(t.GetTwitterTokenRestful))
	ws.Route(ws.GET("/me").To(t.GetMeDetailRestful))
	ws.Route(ws.GET("/follow").To(t.GetFollowerRestful))
	ws.Route(ws.GET("/followids").To(t.GetFollowerIDsRestful))
	ws.Route(ws.GET("/time").To(t.GetTimeLineRestful))
	return ws
}

func (t TwitterResource) RedirectUserToTwitterRestful(request *restful.Request, response *restful.Response) {
	var conf = config.KuberConfig
	fmt.Println(conf)

	fmt.Println("Enter redirect to Twitter")

	redirecttUrl := conf.TwitterRedirect
	if redirecttUrl == "" {
		scheme := "https"
		/*
			if request.Request.TLS != nil {
				scheme = "https"
			}
		*/
		host := request.Request.Host
		if host == "" {
			host = "localhost:9091"
			scheme = "http"
		}
		redirecttUrl = scheme + "://" + host + "/api/v1/twitter/callback"
	}
	fmt.Println("Callback URL=", redirecttUrl)

	requestUrl := TwitterClient.GetAuthURL(redirecttUrl)
	fmt.Println("Request URL: " + requestUrl)

	http.Redirect(response.ResponseWriter, request.Request, requestUrl, http.StatusTemporaryRedirect)
	fmt.Println("Leaving redirect...")

}

func (t TwitterResource) GetTwitterTokenRestful(request *restful.Request, response *restful.Response) {

	fmt.Println("Enter Twitter Callback.")

	var twitterUser TwitterUser

	values := request.Request.URL.Query()
	fmt.Printf("%v\n", values)
	tokenKey := values.Get("oauth_token")
	verificationCode := values.Get("oauth_verifier")

	session, _ := Store.Get(request.Request, SessionName)
	value := session.Values[SessionKey]

	if value == nil {
		json.Unmarshal([]byte(`
				{"username": "", "isAuth": false},
			`), &twitterUser)
	} else {

		fmt.Println(tokenKey)
		fmt.Println(verificationCode)

		twitterUser, _ = value.(TwitterUser)
		fmt.Println(twitterUser)

		TwitterClient.CompleteAuth(tokenKey, verificationCode)

		twitterUser.Oauth_token = tokenKey
		twitterUser.Oauth_verifier = verificationCode
		twitterUser.IsAuth = true

	}

	fmt.Println(twitterUser)
	session.Values[SessionKey] = twitterUser
	session.Save(request.Request, response.ResponseWriter)

	Save(tokenKey, verificationCode)

	redirectURL := fmt.Sprintf(cff.KuberPlane+"#/auth/twitter/callback?token=%s&code=%s", tokenKey, verificationCode)
	http.Redirect(response.ResponseWriter, request.Request, redirectURL, http.StatusTemporaryRedirect)

}

func (t TwitterResource) GetMeDetailRestful(request *restful.Request, response *restful.Response) {
	me, _, _ := TwitterClient.VerifyCredentials()
	fmt.Println("Me Detail =", me)
	response.WriteEntity(me)
}

func (t TwitterResource) GetTimeLineRestful(request *restful.Request, response *restful.Response) {
	GetTimeLine(response.ResponseWriter, request.Request)
}

func (t TwitterResource) GetFollowerRestful(request *restful.Request, response *restful.Response) {
	GetFollower(response.ResponseWriter, request.Request)
}

func (t TwitterResource) GetFollowerIDsRestful(request *restful.Request, response *restful.Response) {
	GetFollowerIDs(response.ResponseWriter, request.Request)
}

func isAuth(w http.ResponseWriter, r *http.Request) bool {
	session, _ := Store.Get(r, SessionName)
	value := session.Values[SessionKey]
	fmt.Println(value)
	var twitterUser TwitterUser
	if value == nil {
		json.Unmarshal([]byte(`
			{"username": "", "isAuth": false},
		`), &twitterUser)
	} else {
		twitterUser, _ = value.(TwitterUser)
	}
	fmt.Println(twitterUser)
	session.Values[SessionKey] = twitterUser
	err := session.Save(r, w)
	if err != nil {
		log.Println(err)
	}
	return twitterUser.IsAuth
}
