package twitter

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/datalayer/kuber/config"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/sessions"
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

var twitterSession *TwitterSession

func (t TwitterResource) WebService() *restful.WebService {

	gob.Register(TwitterUser{})
	twitterSession = NewTwitterSession(config.KuberConfig.TwitterConsumerKey, config.KuberConfig.TwitterConsumerSecret)

	Store.Options = &sessions.Options{
		Path:     "/",      // to match all requests
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: false,
	}

	ws := new(restful.WebService)
	ws.Path("/api/v1/twitter").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(t.RedirecToTwitter))
	ws.Route(ws.GET("/callback").To(t.GetTwitterToken))
	ws.Route(ws.POST("/me").To(t.GetMe))
	/*
		ws.Route(ws.GET("/follow").To(t.GetFollower))
		ws.Route(ws.GET("/followids").To(t.GetFollowerIDs))
		ws.Route(ws.GET("/time").To(t.GetTimeLine))
	*/
	return ws
}

func (t TwitterResource) RedirecToTwitter(request *restful.Request, response *restful.Response) {
	var conf = config.KuberConfig
	fmt.Println(conf)

	fmt.Println("Enter redirect to Twitter")

	redirecttUrl := conf.TwitterRedirect
	if redirecttUrl == "" {
		scheme := "https"
		host := request.Request.Host
		if strings.HasPrefix(host, "localhost") {
			scheme = "http"
		}
		redirecttUrl = scheme + "://" + host + "/api/v1/twitter/callback"
	}
	fmt.Println("Callback URL=", redirecttUrl)

	requestUrl := twitterSession.GetAuthURL(redirecttUrl)
	fmt.Println("Request URL: " + requestUrl)

	http.Redirect(response.ResponseWriter, request.Request, requestUrl, http.StatusTemporaryRedirect)
	fmt.Println("Leaving redirect...")

}

func (t TwitterResource) GetTwitterToken(request *restful.Request, response *restful.Response) {

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

		twitterSession.CompleteAuth(tokenKey, verificationCode)

		twitterUser.Oauth_token = tokenKey
		twitterUser.Oauth_verifier = verificationCode
		twitterUser.IsAuth = true

	}

	fmt.Println(twitterUser)
	session.Values[SessionKey] = twitterUser
	session.Save(request.Request, response.ResponseWriter)

	Save(tokenKey, verificationCode)

	redirectURL := fmt.Sprintf(config.KuberConfig.KuberPlane+"/#/auth/twitter/callback?token=%s&code=%s", tokenKey, verificationCode)
	http.Redirect(response.ResponseWriter, request.Request, redirectURL, http.StatusTemporaryRedirect)

}

func (t TwitterResource) GetMe(request *restful.Request, response *restful.Response) {
	me, _, _ := twitterSession.VerifyCredentials()
	fmt.Println("Me Detail =", me)
	response.WriteEntity(me)
}

/*
func (t TwitterResource) GetTimeLine(request *restful.Request, response *restful.Response) {
	GetTimeLine(response.ResponseWriter, request.Request)
}

func (t TwitterResource) GetFollower(request *restful.Request, response *restful.Response) {
	GetFollower(response.ResponseWriter, request.Request)
}

func (t TwitterResource) GetFollowerIDs(request *restful.Request, response *restful.Response) {
	GetFollowerIDs(response.ResponseWriter, request.Request)
}
*/
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
