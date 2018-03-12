package rest

import (
	"github.com/datalayer/kuber/cloud"
	"github.com/datalayer/kuber/config"
	"github.com/datalayer/kuber/helm"
	k "github.com/datalayer/kuber/k8s"
	"github.com/datalayer/kuber/user"
	wso "github.com/datalayer/kuber/ws"
	restful "github.com/emicklei/go-restful"
)

func SetupGoRestful(wsContainer *restful.Container) {

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  AllowedHeaders(),
		AllowedHeaders: AllowedHeaders(),
		AllowedMethods: AllowedMethods(),
		AllowedDomains: AllowedOrigins(),
		CookiesAllowed: true,
		Container:      wsContainer,
	}
	wsContainer.Filter(cors.Filter)
	// Add container filter to respond to OPTIONS.
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	// Web Socket Resources.
	wsoc := wso.WsResource{wso.WsMessage{}}
	wsContainer.Add(wsoc.WebService())

	// Config Resources.
	conf := config.ConfigResource{config.Config{}}
	wsContainer.Add(conf.WebService())

	// Cloud Resources.
	clo := cloud.CloudResource{}
	wsContainer.Add(clo.WebService())

	// K8S Resources.
	k := k.ClusterResource{}
	wsContainer.Add(k.WebService())

	// Helm Resources.
	h := helm.HelmResource{}
	wsContainer.Add(h.WebService())

	// User Resources.
	u := user.UserResource{map[string]user.User{}}
	wsContainer.Add(u.WebService())

}