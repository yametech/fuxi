package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"k8s.io/api/extensions/v1beta1"
	"net/http"
)

// Get Ingress
func (w *WorkloadsAPI) GetIngress(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.ingress.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Ingress
func (w *WorkloadsAPI) ListIngress(g *gin.Context) {
	list, err := w.ingress.List("", "", 0, 10000, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	ingressList := &v1beta1.IngressList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, ingressList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, ingressList)
}
