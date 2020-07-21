package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	corev1 "k8s.io/api/core/v1"
)

type patchNodeAnnotation struct {
	Node string `json:"node"`
	Host string `json:"host"`
	Rack string `json:"rack"`
	Zone string `json:"zone"`
}

func (w *WorkloadsAPI) GeoAnnotateNode(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	pad := patchNodeAnnotation{}
	err = json.Unmarshal(rawData, &pad)
	if err != nil {
		common.ToInternalServerError(g, rawData, err)
		return
	}
	patchData := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]string{
				"nuwa.kubernetes.io/host": pad.Host,
				"nuwa.kubernetes.io/rack": pad.Rack,
				"nuwa.kubernetes.io/zone": pad.Zone,
			},
		},
	}
	newObj, err := w.node.Patch("", pad.Node, patchData)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}

// Get Node
func (w *WorkloadsAPI) GetNode(g *gin.Context) {
	node := g.Param("node")
	item, err := w.node.Get("", node)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
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
