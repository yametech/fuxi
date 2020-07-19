package handler

import (
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get IP
func (w *WorkloadsAPI) GetIP(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.ip.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List IP
func (w *WorkloadsAPI) ListIP(g *gin.Context) {
	list, err := resourceList(g, w.ip)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
