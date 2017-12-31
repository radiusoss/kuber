package config

import (
	"encoding/gob"
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
)

var sessionName = "kuber-config"
var sessionKey = "config-state"

var store = sessions.NewCookieStore([]byte("spl-session-secret"))

func init() {
	gob.Register(Config{})
	store.Options = &sessions.Options{
		Path:     "/",      // to match all requests
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: false,
	}
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(KuberConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, bytes)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
