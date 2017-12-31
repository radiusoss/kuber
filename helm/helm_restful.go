package helm

import (
	"github.com/datalayer/kuber/cluster"

	"github.com/emicklei/go-restful"
	c "github.com/kris-nova/kubicorn/apis/cluster"
)

const (
	clusterName = "kuber"
)

type HelmResource struct {
}

func (h HelmResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/helm").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("").To(h.GetDeployments))
	ws.Route(ws.GET("/{releaseName}").To(h.GetDeployment))
	ws.Route(ws.POST("/{chartName}/{releaseName}").To(h.Deploy))
	return ws
}

func (h HelmResource) Deploy(request *restful.Request, response *restful.Response) {
	res, _ := Deploy(getCluster(), request.PathParameter("chartName"), request.PathParameter("releaseName"), nil)
	response.WriteEntity(res)
}

func (h HelmResource) GetDeployment(request *restful.Request, response *restful.Response) {
	deployment, _ := GetDeployment(getCluster(), request.PathParameter("releaseName"))
	response.WriteEntity(deployment)
}

func (h HelmResource) GetDeployments(request *restful.Request, response *restful.Response) {
	deployments, _ := GetDeployments(getCluster(), nil)
	response.WriteEntity(deployments)
}

func getCluster() *c.Cluster {
	kc := cluster.KuberCluster{
		Name: clusterName,
	}
	cluster, err := cluster.GetCluster(kc)
	if err != nil {
		panic(err.Error())
	}
	return cluster
}
