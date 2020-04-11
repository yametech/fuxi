package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	corev1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Event
func (w *WorkloadsAPI) GetEvent(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.event.Get(dyn.ResourceEvent, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Event
func (w *WorkloadsAPI) ListEvent(g *gin.Context) {
	list, _ := w.event.List(dyn.ResourceEvent, "", "", 0, 10000, nil)
	eventList := &corev1.EventList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, eventList)
	g.JSON(http.StatusOK, eventList)
}
