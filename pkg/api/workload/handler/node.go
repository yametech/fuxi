package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	corev1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Node
func (w *WorkloadsAPI) GetNode(g *gin.Context) {
	node := g.Param("node")
	item, err := w.node.Get(dyn.ResourceNode, "", node)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Node
func (w *WorkloadsAPI) ListNode(g *gin.Context) {
	list, _ := w.node.List(dyn.ResourceNode, "", "", 0, 10000, nil)
	nodeList := &corev1.NodeList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, nodeList)
	g.JSON(http.StatusOK, nodeList)
}
