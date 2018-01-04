package microsoft

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
	"github.com/spf13/viper"
)

type MicrosoftResource struct {
}

var cff config.Config

func (m MicrosoftResource) WebService(cf config.Config) *restful.WebService {

	cff = cf

	err := viper.Unmarshal(&config.KuberConfig)
	if err != nil {
		log.Error("Unable to decode into struct, %v", err)
	}
	log.Info("Kuber Config:", config.KuberConfig)

	ws := new(restful.WebService)
	ws.Path("/api/v1/microsoft")
	ws.Route(ws.GET("/callback").To(m.MicrosoftCallback))
	return ws
}

func (m MicrosoftResource) MicrosoftCallback(request *restful.Request, response *restful.Response) {

	fmt.Println("Enter Microsoft Callback.")

	codes, ok := request.Request.URL.Query()["code"]

	if !ok || len(codes) < 1 {
		log.Error("Issue while getting code from Microsoft.")
	} else {
		m.getAuthorizationCode(request, response, codes[0])
	}

}

func (m MicrosoftResource) getAuthorizationCode(request *restful.Request, response *restful.Response, code string) {

	hc := http.Client{}

	form := url.Values{}
	form.Add("code", code)
	form.Add("grant_type", "authorization_code")
	form.Add("client_id", cff.MicrosoftApplicationId)
	form.Add("client_secret", cff.MicrosoftSecret)
	form.Add("scope", cff.MicrosoftScope)
	form.Add("redirect_uri", cff.MicrosoftRedirect)
	log.Info("Form: %v", form)

	req, _ := http.NewRequest("POST", "https://login.microsoftonline.com/common/oauth2/v2.0/token", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	resp, err := hc.Do(req)

	if err != nil {
		log.Error("Unable to decode into struct, %v", err)
	}
	/*
		{
			"access_token": " eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6Ik5HVEZ2ZEstZnl0aEV1THdqcHdBSk9NOW4tQSJ9.eyJhdWQiOiJodHRwczovL3NlcnZpY2UuY29udG9zby5jb20vIiwiaXNzIjoiaHR0cHM6Ly9zdHMud2luZG93cy5uZXQvN2ZlODE0NDctZGE1Ny00Mzg1LWJlY2ItNmRlNTdmMjE0NzdlLyIsImlhdCI6MTM4ODQ0MDg2MywibmJmIjoxMzg4NDQwODYzLCJleHAiOjEzODg0NDQ3NjMsInZlciI6IjEuMCIsInRpZCI6IjdmZTgxNDQ3LWRhNTctNDM4NS1iZWNiLTZkZTU3ZjIxNDc3ZSIsIm9pZCI6IjY4Mzg5YWUyLTYyZmEtNGIxOC05MWZlLTUzZGQxMDlkNzRmNSIsInVwbiI6ImZyYW5rbUBjb250b3NvLmNvbSIsInVuaXF1ZV9uYW1lIjoiZnJhbmttQGNvbnRvc28uY29tIiwic3ViIjoiZGVOcUlqOUlPRTlQV0pXYkhzZnRYdDJFYWJQVmwwQ2o4UUFtZWZSTFY5OCIsImZhbWlseV9uYW1lIjoiTWlsbGVyIiwiZ2l2ZW5fbmFtZSI6IkZyYW5rIiwiYXBwaWQiOiIyZDRkMTFhMi1mODE0LTQ2YTctODkwYS0yNzRhNzJhNzMwOWUiLCJhcHBpZGFjciI6IjAiLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJhY3IiOiIxIn0.JZw8jC0gptZxVC-7l5sFkdnJgP3_tRjeQEPgUn28XctVe3QqmheLZw7QVZDPCyGycDWBaqy7FLpSekET_BftDkewRhyHk9FW_KeEz0ch2c3i08NGNDbr6XYGVayNuSesYk5Aw_p3ICRlUV1bqEwk-Jkzs9EEkQg4hbefqJS6yS1HoV_2EsEhpd_wCQpxK89WPs3hLYZETRJtG5kvCCEOvSHXmDE6eTHGTnEgsIk--UlPe275Dvou4gEAwLofhLDQbMSjnlV5VLsjimNBVcSRFShoxmQwBJR_b2011Y5IuD6St5zPnzruBbZYkGNurQK63TJPWmRd3mbJsGM0mf3CUQ",
			"token_type": "Bearer",
			"expires_in": "3600",
			"expires_on": "1388444763",
			"resource": "https://service.contoso.com/",
			"refresh_token": "AwABAAAAvPM1KaPlrEqdFSBzjqfTGAMxZGUTdM0t4B4rTfgV29ghDOHRc2B-C_hHeJaJICqjZ3mY2b_YNqmf9SoAylD1PycGCB90xzZeEDg6oBzOIPfYsbDWNf621pKo2Q3GGTHYlmNfwoc-OlrxK69hkha2CF12azM_NYhgO668yfcUl4VBbiSHZyd1NVZG5QTIOcbObu3qnLutbpadZGAxqjIbMkQ2bQS09fTrjMBtDE3D6kSMIodpCecoANon9b0LATkpitimVCrl-NyfN3oyG4ZCWu18M9-vEou4Sq-1oMDzExgAf61noxzkNiaTecM-Ve5cq6wHqYQjfV9DOz4lbceuYCAA",
			"scope": "https%3A%2F%2Fgraph.microsoft.com%2Fmail.read",
			"id_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJhdWQiOiIyZDRkMTFhMi1mODE0LTQ2YTctODkwYS0yNzRhNzJhNzMwOWUiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC83ZmU4MTQ0Ny1kYTU3LTQzODUtYmVjYi02ZGU1N2YyMTQ3N2UvIiwiaWF0IjoxMzg4NDQwODYzLCJuYmYiOjEzODg0NDA4NjMsImV4cCI6MTM4ODQ0NDc2MywidmVyIjoiMS4wIiwidGlkIjoiN2ZlODE0NDctZGE1Ny00Mzg1LWJlY2ItNmRlNTdmMjE0NzdlIiwib2lkIjoiNjgzODlhZTItNjJmYS00YjE4LTkxZmUtNTNkZDEwOWQ3NGY1IiwidXBuIjoiZnJhbmttQGNvbnRvc28uY29tIiwidW5pcXVlX25hbWUiOiJmcmFua21AY29udG9zby5jb20iLCJzdWIiOiJKV3ZZZENXUGhobHBTMVpzZjd5WVV4U2hVd3RVbTV5elBtd18talgzZkhZIiwiZmFtaWx5X25hbWUiOiJNaWxsZXIiLCJnaXZlbl9uYW1lIjoiRnJhbmsifQ."
		}
	*/
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

	u := cff.KuberPlane + "#/auth/microsoft/callback" + "?access_token=" + data.AccessToken
	fmt.Println(u)
	http.Redirect(response.ResponseWriter, request.Request, u, http.StatusTemporaryRedirect)

}
