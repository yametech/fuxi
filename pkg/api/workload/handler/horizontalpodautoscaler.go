package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	autoscalingv2beta1 "k8s.io/api/autoscaling/v2beta1"
	"net/http"
)

// Get HorizontalPodAutoscaler
func (w *WorkloadsAPI) GetHorizontalPodAutoscaler(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.horizontalPodAutoscaler.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List HorizontalPodAutoscaler
func (w *WorkloadsAPI) ListHorizontalPodAutoscaler(g *gin.Context) {
	list, err := w.horizontalPodAutoscaler.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	horizontalPodAutoscalerList := &autoscalingv2beta1.HorizontalPodAutoscalerList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, horizontalPodAutoscalerList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, horizontalPodAutoscalerList)
}
