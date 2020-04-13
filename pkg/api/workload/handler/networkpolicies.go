package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/networking/v1"
	"net/http"
)

// Get NetworkPolicy
func (w *WorkloadsAPI) GetNetworkPolicy(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.networkPolicy.Get(dyn.ResourceNetworkPolicy, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List NetworkPolicy
func (w *WorkloadsAPI) ListNetworkPolicy(g *gin.Context) {
	list, _ := w.service.List(dyn.ResourceNetworkPolicy, "", "", 0, 10000, nil)
	networkPolicyList := &v1.NetworkPolicyList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, networkPolicyList)
	g.JSON(http.StatusOK, networkPolicyList)
}
