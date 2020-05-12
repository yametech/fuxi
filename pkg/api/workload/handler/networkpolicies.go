package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/networking/v1"
	"net/http"
)

// Get NetworkPolicy
func (w *WorkloadsAPI) GetNetworkPolicy(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.networkPolicy.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List NetworkPolicy
func (w *WorkloadsAPI) ListNetworkPolicy(g *gin.Context) {
	list, err := w.networkPolicy.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	networkPolicyList := &v1.NetworkPolicyList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, networkPolicyList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, networkPolicyList)
}
