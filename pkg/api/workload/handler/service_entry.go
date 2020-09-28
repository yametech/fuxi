package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
)

func (w *WorkloadsAPI) GetServiceEntry(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.gateway.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

func (w *WorkloadsAPI) ListServiceEntry(g *gin.Context) {
	list, err := resourceList(g, w.gateway)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
