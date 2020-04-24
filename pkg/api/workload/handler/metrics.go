package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *WorkloadsAPI) Metrics2(g *gin.Context) {
	body, err := g.GetRawData()
	if err != nil {
		panic(err)
	}
	newParams := make(map[string]string)
	newParams["query"] = g.Query("query")
	newParams["start"] = g.Query("start")
	newParams["end"] = g.Query("end")
	newParams["step"] = g.Query("step")
	newParams["kubernetes_namespace"] = g.Query("kubernetes_namespace")

	bufRaw := bytes.NewBuffer(nil)
	if err := w.metrics.PostProxyToPrometheus(newParams, string(body), bufRaw); err != nil {
		//panic(err)
	}

	g.JSON(http.StatusOK, bufRaw.String())
}

func (w *WorkloadsAPI) Metrics(g *gin.Context) {
	metricsData, err := w.metrics.PostMetrics(g.Request.URL.RawQuery)
	if err != nil {
		g.JSON(http.StatusInternalServerError,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "backend service get error"},
		)
		return
	}
	g.JSON(http.StatusOK, metricsData)
}
