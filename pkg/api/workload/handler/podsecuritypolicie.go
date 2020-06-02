package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"k8s.io/kubernetes/pkg/apis/policy"
	"net/http"
)

// Get PodSecurityPolicie
func (w *WorkloadsAPI) GetPodSecurityPolicie(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.podsecuritypolicies.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List PodSecurityPolicie
func (w *WorkloadsAPI) ListPodSecurityPolicie(g *gin.Context) {
	list, err := w.podsecuritypolicies.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	podSecurityPolicyList := &policy.PodSecurityPolicyList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, podSecurityPolicyList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, podSecurityPolicyList)
}
