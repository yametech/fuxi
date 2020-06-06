package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"
)

// Get PodSecurityPolicie
func (w *WorkloadsAPI) GetPodSecurityPolicie(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.podsecuritypolicies.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List PodSecurityPolicie
func (w *WorkloadsAPI) ListPodSecurityPolicie(g *gin.Context) {
	list, err := w.podsecuritypolicies.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
