package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/workload/template"
	//dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	//workload "github.com/yametech/fuxi/pkg/service/workload"
	"net/http"
)

func (w *WorkloadsAPI) GetDeployment(g *gin.Context) {

}

func (w *WorkloadsAPI) ListDeployment(g *gin.Context) {

}

func (w *WorkloadsAPI) ApplyDeployment(g *gin.Context) {
	deploymentRequest := &template.DeploymentRequest{}
	if err := g.ShouldBind(deploymentRequest); err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter",
			},
		)
		return
	}
	//deploymentService := workload.NewDeployment()

	//deploymentService.Apply(dyn.Deployment,)
}

func (w *WorkloadsAPI) DeleteDeployment(g *gin.Context) {

}
