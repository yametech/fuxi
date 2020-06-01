package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"net/http"
)

// GetDeploymentScale get deployment scale
func (w *WorkloadsAPI) GetDeploymentScale(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")

	w.generic.SetGroupVersionResource(types.ResourceDeployment)
	item, err := w.generic.RemoteGet(namespace, name, "scale")
	if err != nil {
		common.ToInternalServerError(g, "", err)
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
		common.ToRequestParamsError(g, err)
		return
	}
	pathData := make(map[string]interface{})
	err = json.Unmarshal(rawData, &pathData)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	w.generic.SetGroupVersionResource(types.ResourceDeployment)
	_, err = w.generic.Patch(namespace, name, pathData)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, "")
}
