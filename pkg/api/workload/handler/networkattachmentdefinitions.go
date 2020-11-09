package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"
)

// Get NetworkPolicy
func (w *WorkloadsAPI) GetNetworkAttachmentDefinition(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.networkAttachmentDefinitions.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List NetworkPolicy
func (w *WorkloadsAPI) ListNetworkAttachmentDefinition(g *gin.Context) {
	list, err := resourceList(g, w.networkAttachmentDefinitions)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	networkAttachmentDefinitionList := &v1.NetworkAttachmentDefinition{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, networkAttachmentDefinitionList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, networkAttachmentDefinitionList)
}
