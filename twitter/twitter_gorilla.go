package twitter

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/datalayer/kuber/config"
	"github.com/gorilla/sessions"
)

func init() {

	gob.Register(TwitterUser{})

	var conf = config.KuberConfig
	TwitterClient = NewServerClient(conf.TwitterConsumerKey, conf.TwitterConsumerSecret)

	Store.Options = &sessions.Options{
		Path:     "/",      // to match all requests
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: false,
	}

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

		http.Redirect(w, r, conf.TwitterRedirect, http.StatusTemporaryRedirect)

	}
}

func RedirectUserToTwitter(w http.ResponseWriter, r *http.Request) {

	var conf = config.KuberConfig
	fmt.Println(conf)

	fmt.Println("Enter redirect to twitter")

	redirecttUrl := conf.TwitterRedirect
	if redirecttUrl == "" {
		redirecttUrl = r.URL.String()
		strings.Replace(redirecttUrl, "", "", 1)
	}
	fmt.Println("Callback URL=", redirecttUrl)

	requestUrl := TwitterClient.GetAuthURL(redirecttUrl)
	fmt.Println("Request URL: " + requestUrl)

	http.Redirect(w, r, requestUrl, http.StatusTemporaryRedirect)
	fmt.Println("Leaving redirect...")

}

func GetTwitterToken(w http.ResponseWriter, r *http.Request) {

	var twitterUser TwitterUser

	fmt.Println("Enter Get twitter token.")

	values := r.URL.Query()

	fmt.Println(values)

	tokenKey := values.Get("oauth_token")
	verificationCode := values.Get("oauth_verifier")

	session, _ := Store.Get(r, SessionName)
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
	session.Save(r, w)

	Save(tokenKey, verificationCode)

	redirectURL := fmt.Sprintf("http://localhost:4326/#/auth/twitter/callback"+"?token=%s&code=%s", tokenKey, verificationCode)
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
