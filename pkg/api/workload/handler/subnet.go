package handler

import (
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get SubNet
func (w *WorkloadsAPI) GetSubNet(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.subnet.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List SubNet
func (w *WorkloadsAPI) ListSubNet(g *gin.Context) {
	list, err := resourceList(g, w.subnet)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
