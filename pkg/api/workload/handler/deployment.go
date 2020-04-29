package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	appsv1 "k8s.io/api/apps/v1"
	"net/http"
)

func (w *WorkloadsAPI) GetDeployment(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.deployments.Get(dyn.ResourceDeployment, namespace, name)
	if err != nil {
		panic(err)
	}
	g.JSON(http.StatusOK, item)
}

func (w *WorkloadsAPI) ListDeployment(g *gin.Context) {
	list, _ := w.deployments.List(dyn.ResourceDeployment, "", "", 0, 100, nil)
	deploymentList := &appsv1.DeploymentList{}
	data, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(data, deploymentList)
	g.JSON(http.StatusOK, deploymentList)
}
