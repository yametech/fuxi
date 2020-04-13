package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/api/extensions/v1beta1"
	"net/http"
)

// Get Ingress
func (w *WorkloadsAPI) GetIngress(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.ingress.Get(dyn.ResourceIngress, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Ingress
func (w *WorkloadsAPI) ListIngress(g *gin.Context) {
	list, _ := w.ingress.List(dyn.ResourceIngress, "", "", 0, 10000, nil)
	ingressList := &v1beta1.IngressList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, ingressList)
	g.JSON(http.StatusOK, ingressList)
}
