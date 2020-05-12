package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	"net/http"
)

func (w *WorkloadsAPI) GetDeployment(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.deployments.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

func (w *WorkloadsAPI) ListDeployment(g *gin.Context) {
	list, err := w.deployments.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	deploymentList := &appsv1.DeploymentList{}
	data, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(data, deploymentList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, deploymentList)
}
