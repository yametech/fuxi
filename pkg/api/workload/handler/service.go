package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Service
func (w *WorkloadsAPI) GetService(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.service.Get(dyn.ResourceService, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Service
func (w *WorkloadsAPI) ListService(g *gin.Context) {
	list, _ := w.service.List(dyn.ResourceService, "", "", 0, 10000, nil)
	serviceList := &v1.ServiceList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, serviceList)
	g.JSON(http.StatusOK, serviceList)
}
