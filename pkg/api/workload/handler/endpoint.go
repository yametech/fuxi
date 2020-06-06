package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	corev1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Endpoint
func (w *WorkloadsAPI) GetEndpoint(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.endpoint.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Endpoint
func (w *WorkloadsAPI) ListEndpoint(g *gin.Context) {
	list, err := resourceList(g, w.endpoint)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	endpointList := &corev1.EndpointsList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, endpointList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, endpointList)
}
