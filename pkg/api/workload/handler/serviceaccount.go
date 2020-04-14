package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

// Get ServiceAccount
func (w *WorkloadsAPI) GetServiceAccount(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.serviceAccount.Get(dyn.ResourceServiceAccount, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ServiceAccount
func (w *WorkloadsAPI) ListServiceAccount(g *gin.Context) {
	list, _ := w.serviceAccount.List(dyn.ResourceServiceAccount, "", "", 0, 10000, nil)
	serviceAccountList := &v1.ServiceAccountList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, serviceAccountList)
	g.JSON(http.StatusOK, serviceAccountList)
}
