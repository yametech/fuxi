package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
)

func (w *WorkloadsAPI) GetVirtulService(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.virtualService.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

func (w *WorkloadsAPI) ListVirtulService(g *gin.Context) {
	list, err := resourceList(g, w.virtualService)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
