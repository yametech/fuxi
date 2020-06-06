package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

// Get ResourceQuota
func (w *WorkloadsAPI) GetResourceQuota(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.resourceQuota.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ResourceQuota
func (w *WorkloadsAPI) ListResourceQuota(g *gin.Context) {
	list, err := resourceList(g, w.resourceQuota)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	resourceQuotaList := &v1.ResourceQuotaList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, resourceQuotaList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, resourceQuotaList)
}
