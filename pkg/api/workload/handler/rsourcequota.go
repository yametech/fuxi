package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

// Get resourceQuota
func (w *WorkloadsAPI) GetResourceQuota(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.resourceQuota.Get(dyn.ResourceResourceQuota, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List resourceQuota
func (w *WorkloadsAPI) ListResourceQuota(g *gin.Context) {
	list, _ := w.resourceQuota.List(dyn.ResourceResourceQuota, "", "", 0, 10000, nil)
	resourceQuotaList := &v1.ResourceQuotaList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, resourceQuotaList)
	g.JSON(http.StatusOK, resourceQuotaList)
}
