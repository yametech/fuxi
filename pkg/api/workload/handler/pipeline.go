package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"
)

// Get Pipeline
func (w *WorkloadsAPI) GetPipeline(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.pipeline.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Pipeline
func (w *WorkloadsAPI) ListPipeline(g *gin.Context) {
	list, err := resourceList(g, w.pipeline)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
