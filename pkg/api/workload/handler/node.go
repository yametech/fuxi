package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	corev1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Node
func (w *WorkloadsAPI) GetNode(g *gin.Context) {
	node := g.Param("node")
	item, err := w.node.Get("", node)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Node
func (w *WorkloadsAPI) ListNode(g *gin.Context) {
	list, err := w.node.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	nodeList := &corev1.NodeList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, nodeList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, nodeList)
}
