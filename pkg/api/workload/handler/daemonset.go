package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
)

// Get DaemonSet
func (w *WorkloadsAPI) GetDaemonSet(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.daemonSet.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List DaemonSet
func (w *WorkloadsAPI) ListDaemonSet(g *gin.Context) {
	list, err := resourceList(g, w.daemonSet)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	daemonSetList := &v1.DaemonSetList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, daemonSetList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, daemonSetList)
}
