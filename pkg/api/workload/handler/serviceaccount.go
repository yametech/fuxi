package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

// Get ServiceAccount
func (w *WorkloadsAPI) GetServiceAccount(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.serviceAccount.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ServiceAccount
func (w *WorkloadsAPI) ListServiceAccount(g *gin.Context) {
	list, err := resourceList(g, w.serviceAccount)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	serviceAccountList := &v1.ServiceAccountList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, serviceAccountList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, serviceAccountList)
}
