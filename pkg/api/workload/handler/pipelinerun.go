package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"
)

// Get PipelineRun
func (w *WorkloadsAPI) GetPipelineRun(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.pipelineRun.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List PipelineRun
func (w *WorkloadsAPI) ListPipelineRun(g *gin.Context) {
	list, err := resourceList(g, w.pipelineRun)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
