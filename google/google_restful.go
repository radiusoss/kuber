package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/datalayer/kuber/config"
	"github.com/datalayer/kuber/log"
	"github.com/emicklei/go-restful"
)

type GoogleResource struct {
	states map[int]int
}

func (m GoogleResource) WebService() *restful.WebService {
	m.states = make(map[int]int)
	ws := new(restful.WebService)
	ws.Path("/api/v1/google")
	ws.Route(ws.GET("").To(m.Authorize))
	ws.Route(ws.GET("/callback").To(m.Callback))
	return ws
}

func (m GoogleResource) Authorize(request *restful.Request, response *restful.Response) {
	fmt.Println("Enter Google Redirect...")
	values := request.Request.URL.Query()
	fmt.Printf("%v\n", values)
	redirectUrl := getRedirectUrl(request)
	fmt.Println("Callback URL=", redirectUrl)
	values.Add("redirect_uri", redirectUrl)
	//	state := rand.New(time.Now().UnixNano()).Intn(int(MaxUint >> 1))
	//	m.states[state] = state
	u := "https://accounts.google.com/o/oauth2/v2/auth?" + values.Encode()
	fmt.Println("Redirect to:", u)
	http.Redirect(response.ResponseWriter, request.Request, u, http.StatusTemporaryRedirect)
	fmt.Println("Exit Google Redirect...")
}

func (m GoogleResource) Callback(request *restful.Request, response *restful.Response) {

	fmt.Println("Enter Google Callback...")

	codes, ok := request.Request.URL.Query()["code"]

	if !ok || len(codes) < 1 {
		log.Error("Issue while getting code from Google.")
	} else {

		code := codes[0]
		log.Info("code: ", code)

		hc := http.Client{}

		redirecttUrl := getRedirectUrl(request)

		form := url.Values{}
		form.Add("code", code)
		//		form.Add("grant_type", "authorization_code")
		form.Add("client_id", config.KuberConfig.GoogleClientId)
		form.Add("client_secret", config.KuberConfig.GoogleSecret)
		form.Add("scope", config.KuberConfig.GoogleScope)
		form.Add("redirect_uri", redirecttUrl)
		form.Add("grant_type", "authorization_code")
		log.Info("Form: %v", form)

		req, _ := http.NewRequest("POST", "https://www.googleapis.com/oauth2/v4/token", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
		resp, err := hc.Do(req)

		if err != nil {
			log.Error("Unable to decode into struct, %v", err)
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(responseData))

		var data struct {
			AccessToken string `json:"access_token"`
		}
		err = json.Unmarshal([]byte(responseData), &data)
		if err != nil {
			log.Error("Unable to decode into struct, %v", err)
		}
		fmt.Println(data)

		u := config.KuberConfig.KuberBoard

		if u == "" {
			scheme := "https"
			host := request.Request.Host
			if strings.HasPrefix(host, "localhost") {
				scheme = "http"
			}
			u = scheme + "://" + host + "/#/auth/google/callback" + "?access_token=" + data.AccessToken
		} else {
			u = u + "/#/auth/google/callback" + "?access_token=" + data.AccessToken
		}

		fmt.Println("Redirecting after Google callback to:", u)

		http.Redirect(response.ResponseWriter, request.Request, u, http.StatusTemporaryRedirect)

		fmt.Println("Exit Google Callback...")

	}

}

func getRedirectUrl(request *restful.Request) string {
	redirectUrl := config.KuberConfig.GoogleRedirect
	if redirectUrl == "" {
		scheme := "https"
		host := request.Request.Host
		if strings.HasPrefix(host, "localhost") {
			scheme = "http"
		}
		redirectUrl = scheme + "://" + host + "/api/v1/google/callback"
	}
	return redirectUrl
}
