package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	autoscalingv2beta1 "k8s.io/api/autoscaling/v2beta1"
	"net/http"
)

// Get HorizontalPodAutoscaler
func (w *WorkloadsAPI) GetHorizontalPodAutoscaler(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.horizontalPodAutoscaler.Get(dyn.ResourceHorizontalPodAutoscaler, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List HorizontalPodAutoscaler
func (w *WorkloadsAPI) ListHorizontalPodAutoscaler(g *gin.Context) {
	list, _ := w.horizontalPodAutoscaler.List(dyn.ResourceHorizontalPodAutoscaler, "", "", 0, 10000, nil)
	horizontalPodAutoscalerList := &autoscalingv2beta1.HorizontalPodAutoscalerList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, horizontalPodAutoscalerList)
	g.JSON(http.StatusOK, horizontalPodAutoscalerList)
}
