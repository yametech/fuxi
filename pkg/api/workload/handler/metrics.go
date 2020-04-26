package handler

import (
	"github.com/gin-gonic/gin"
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
