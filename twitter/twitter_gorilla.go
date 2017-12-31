package twitter

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/datalayer/kuber/config"
	"github.com/gorilla/sessions"
)

var ConsumerKey string
var ConsumerSecret string

var TwitterClient *ServerClient

var sessionName = "kuber-twitter"
var sessionKey = "twitter-state"

var store = sessions.NewCookieStore([]byte("twitter-session-secret"))

type TwitterUser struct {
	Username       string `json:"username"`
	Oauth_token    string `json:"oauth_token"`
	Oauth_verifier string `json:"oauth_verifier"`
	IsAuth         bool   `json:"isAuth"`
}

func init() {

	ConsumerKey = "Fsy5JzXec7wY5mPPsEdsNkAe4"
	ConsumerSecret = "q0suooaCz17lkiHZZi35OoXfBJrAPRyUBi0AssEppP9YXxBSRz"

	gob.Register(TwitterUser{})

	TwitterClient = NewServerClient(ConsumerKey, ConsumerSecret)
	store.Options = &sessions.Options{
		Path:     "/",      // to match all requests
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: false,
	}

}

func isAuth(w http.ResponseWriter, r *http.Request) bool {

	session, _ := store.Get(r, sessionName)
	value := session.Values[sessionKey]
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
	session.Values[sessionKey] = twitterUser

	err := session.Save(r, w)
	if err != nil {
		log.Println(err)
	}

	return twitterUser.IsAuth

}

func MainProcess(w http.ResponseWriter, r *http.Request) {

	if !isAuth(w, r) {

		fmt.Fprintf(w, `
			<body>
			  <center>
			    <a href='/twitter/request'>
			      <img src='https://g.twimg.com/dev/sites/default/files/images_documentation/sign-in-with-twitter-gray.png'>
			    </a>
			  </center>
			</body>
			`)
		return

	} else {

		var conf = config.KuberConfig
		fmt.Println(conf)

		//Logon, redirect to ...
		//		timelineURL := fmt.Sprintf("http://%s/twitter/time", r.Host)
		//		http.Redirect(w, r, timelineURL, http.StatusTemporaryRedirect)
		http.Redirect(w, r, conf.TwitterRedirect, http.StatusTemporaryRedirect)

	}
}

func RedirectUserToTwitter(w http.ResponseWriter, r *http.Request) {

	var conf = config.KuberConfig
	fmt.Println(conf)

	fmt.Println("Enter redirect to twitter")
	fmt.Println("Callback URL=", conf.TwitterRedirect)
	requestUrl := TwitterClient.GetAuthURL(conf.TwitterRedirect)

	http.Redirect(w, r, requestUrl, http.StatusTemporaryRedirect)
	fmt.Println("Leave redirect")
}

func GetTwitterToken(w http.ResponseWriter, r *http.Request) {

	var twitterUser TwitterUser

	fmt.Println("Enter Get twitter token.")

	values := r.URL.Query()

	tokenKey := values.Get("oauth_token")
	verificationCode := values.Get("oauth_verifier")

	session, _ := store.Get(r, sessionName)
	value := session.Values[sessionKey]

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
	session.Values[sessionKey] = twitterUser
	session.Save(r, w)

	Save(tokenKey, verificationCode)

	//	timelineURL := fmt.Sprintf("http://%s/twitter/time", r.Host)
	//	http.Redirect(w, r, timelineURL, http.StatusTemporaryRedirect)

	redirectURL := fmt.Sprintf("http://localhost:4322/#/twitter/auth?token=%s&code=%s", tokenKey, verificationCode)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)

}

func GetTimeLine(w http.ResponseWriter, r *http.Request) {
	timeline, bits, _ := TwitterClient.QueryTimeLine(1)
	fmt.Println("TimeLine=", timeline)
	fmt.Fprintf(w, "The item is: "+string(bits))
}

func GetFollower(w http.ResponseWriter, r *http.Request) {
	followers, bits, _ := TwitterClient.QueryFollower(10)
	fmt.Println("Followers=", followers)
	fmt.Fprintf(w, "The item is: "+string(bits))
}

func GetFollowerIDs(w http.ResponseWriter, r *http.Request) {
	followers, bits, _ := TwitterClient.QueryFollowerIDs(10)
	fmt.Println("Follower IDs=", followers)
	fmt.Fprintf(w, "The item is: "+string(bits))
}

func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	followers, bits, _ := TwitterClient.QueryFollowerById(2244994945)
	fmt.Println("Follower Detail of =", followers)
	fmt.Fprintf(w, "The item is: "+string(bits))
}
