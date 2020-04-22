package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	corev1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Namespace
func (w *WorkloadsAPI) GetNamespace(g *gin.Context) {
	item, err := w.namespace.Get(dyn.ResourceNamespace, "", "")
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Namespaces
func (w *WorkloadsAPI) ListNamespace(g *gin.Context) {
	list, _ := w.namespace.List(dyn.ResourceNamespace, "", "", 0, 10000, nil)
	namespaceList := &corev1.NamespaceList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, namespaceList)
	g.JSON(http.StatusOK, namespaceList)
}
