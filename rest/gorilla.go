package rest

import (
	"net/http"

	"github.com/datalayer/kuber/config"
	"github.com/datalayer/kuber/spl"
	"github.com/datalayer/kuber/twitter"
	"github.com/datalayer/kuber/ws"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupGorilla() *mux.Router {

	r := mux.NewRouter().StrictSlash(false)

	s := r.PathPrefix("/spl/v1").Subrouter()
	s.Methods("GET").Path("/spl").HandlerFunc(spl.GetAllSpl)
	s.Methods("POST").Path("/spl").HandlerFunc(spl.SaveSpl)
	s.Methods("GET").Path("/spl/{name}").HandlerFunc(spl.GetSpl)
	s.Methods("PUT").Path("/spl/{name}").HandlerFunc(spl.UpdateSpl)
	s.Methods("DELETE").Path("/spl/{name}").HandlerFunc(spl.DeleteSpl)

	tw := r.PathPrefix("/twitter").Subrouter()
	tw.Methods("GET").Path("/").HandlerFunc(twitter.MainProcess)
	tw.Methods("GET").Path("/maketoken").HandlerFunc(twitter.GetTwitterToken)
	tw.Methods("GET").Path("/request").HandlerFunc(twitter.RedirectUserToTwitter)
	tw.Methods("GET").Path("/follow").HandlerFunc(twitter.GetFollower)
	tw.Methods("GET").Path("/followids").HandlerFunc(twitter.GetFollowerIDs)
	tw.Methods("GET").Path("/time").HandlerFunc(twitter.GetTimeLine)

	c := r.PathPrefix("/api").Subrouter()
	c.Methods("GET").Path("/config").HandlerFunc(config.GetConfig)

	r.PathPrefix("/echo").HandlerFunc(ws.Echo)
	r.PathPrefix("/pipe").HandlerFunc(ws.Pipe)
	r.PathPrefix("/ws").HandlerFunc(ws.Ws)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./_static/")))

	return r

}

func CredentialsOk() handlers.CORSOption {
	return handlers.AllowCredentials()
}

func HeadersOk() handlers.CORSOption {
	return handlers.AllowedHeaders(AllowedHeaders())
}

func OriginsOk() handlers.CORSOption {
	return handlers.AllowedOrigins(AllowedOrigins())
}

func MethodsOk() handlers.CORSOption {
	return handlers.AllowedMethods(AllowedMethods())

}
