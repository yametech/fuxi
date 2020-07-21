package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "k8s.io/api/networking/v1"
	"net/http"
)

// Get NetworkPolicy
func (w *WorkloadsAPI) GetNetworkPolicy(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.networkPolicy.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List NetworkPolicy
func (w *WorkloadsAPI) ListNetworkPolicy(g *gin.Context) {
	list, err := resourceList(g, w.networkPolicy)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	networkPolicyList := &v1.NetworkPolicyList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, networkPolicyList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, networkPolicyList)
}
