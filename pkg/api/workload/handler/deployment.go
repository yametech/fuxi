package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	appsv1 "k8s.io/api/apps/v1"
	"net/http"
)

func (w *WorkloadsAPI) GetDeployment(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.deployments.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

func (w *WorkloadsAPI) ListDeployment(g *gin.Context) {
	namespace := g.Param("namespace")
	list, err := w.deployments.List(namespace, "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	deploymentList := &appsv1.DeploymentList{}
	data, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(data, deploymentList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, deploymentList)
}
