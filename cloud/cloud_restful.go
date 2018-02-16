package cloud

import (
	"github.com/datalayer/kuber/aws"
	"github.com/emicklei/go-restful"
)

type Cloud struct {
	ID   string `json:"id" description:"identifier of the cloud"`
	Name string `json:"name" description:"name of the cloud"`
}
type CloudResource struct {
}

func (cl CloudResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/kuber/api/v1/cloud").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/aws/{region}/volumes").To(cl.GetAwsVolumes))
	return ws
}

func (cl CloudResource) GetAwsVolumes(request *restful.Request, response *restful.Response) {
	volumes := aws.GetVolumes(request.PathParameter("region"))
	response.WriteEntity(volumes)
}
