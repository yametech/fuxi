package handler

import (
	"github.com/gin-gonic/gin"
	"k8s.io/metrics/pkg/apis/metrics"
	"net/http"
)

func (w *WorkloadsAPI) Metrics(g *gin.Context) {
	body, err := g.GetRawData()
	if err != nil {
		g.JSON(
			http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter",
			},
		)
	}
	newParams := make(map[string]string)
	newParams["start"] = g.Query("start")
	newParams["end"] = g.Query("end")
	newParams["step"] = g.Query("step")
	newParams["kubernetes_namespace"] = g.Query("kubernetes_namespace")

	bufRaw, err := w.metrics.ProxyToPrometheus(newParams, body)
	if err != nil {
		g.JSON(
			http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   "",
				msg:    err.Error(),
				status: "backend service get error",
			},
		)
		return
	}
	g.JSON(http.StatusOK, bufRaw)
}

func (w *WorkloadsAPI) NodeMetrics(g *gin.Context) {
	nodeMetricsList := &metrics.NodeMetricsList{}
	if err := w.metrics.GetNodeMetricsList(nodeMetricsList); err != nil {
		g.JSON(
			http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   "",
				msg:    err.Error(),
				status: "backend service get error",
			},
		)
		return
	}
	g.JSON(http.StatusOK, nodeMetricsList)
}

func (w *WorkloadsAPI) PodMetrics(g *gin.Context) {
	namespace := g.Query("namespace")
	name := g.Query("name")
	podMetrics := &metrics.PodMetrics{}
	if err := w.metrics.GetPodMetrics(namespace, name, podMetrics); err != nil {
		g.JSON(
			http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   "",
				msg:    err.Error(),
				status: "backend service get error",
			},
		)
		return
	}
	g.JSON(http.StatusOK, podMetrics)
}

func (w *WorkloadsAPI) PodMetricsList(g *gin.Context) {
	namespace := g.Query("namespace")
	podMetricsList := &metrics.PodMetricsList{}
	if err := w.metrics.GetPodMetricsList(namespace, podMetricsList); err != nil {
		g.JSON(
			http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   "",
				msg:    err.Error(),
				status: "backend service get error",
			},
		)
		return
	}
	g.JSON(http.StatusOK, podMetricsList)
}
