package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Service
func (w *WorkloadsAPI) GetService(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.service.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Service
func (w *WorkloadsAPI) ListService(g *gin.Context) {
	list, err := resourceList(g, w.service)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	serviceList := &v1.ServiceList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, serviceList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, serviceList)
}
