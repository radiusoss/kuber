package cluster

import (
	"github.com/emicklei/go-restful"
)

type Cluster struct {
	ID   string `json:"id" description:"identifier of the kuber"`
	Name string `json:"name" description:"name of the cluster"`
}
type ClusterResource struct {
}

func (cl ClusterResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/cluster").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	return ws
}
