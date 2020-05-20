package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"net/http"
)

// GetDeploymentScale get deployment scale
func (w *WorkloadsAPI) GetDeploymentScale(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")

	w.generic.SetGroupVersionResource(dyn.ResourceDeployment)
	item, err := w.generic.RemoteGet(namespace, name, "scale")
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// PutDeploymentScale update scale deployments replicas
func (w *WorkloadsAPI) PutDeploymentScale(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")

	rawData, err := g.GetRawData()
	if err != nil {
		toRequestParamsError(g, err)
		return
	}
	pathData := make(map[string]interface{})
	err = json.Unmarshal(rawData, &pathData)
	if err != nil {
		toRequestParamsError(g, err)
		return
	}

	w.generic.SetGroupVersionResource(dyn.ResourceDeployment)
	_, err = w.generic.Path(namespace, name, pathData)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, "")
}
