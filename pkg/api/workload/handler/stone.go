package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"
)

// Get Stone
func (w *WorkloadsAPI) GetStone(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")

	item, err := w.stone.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List stone
func (w *WorkloadsAPI) ListStone(g *gin.Context) {
	list, err := resourceList(g, w.stone)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
