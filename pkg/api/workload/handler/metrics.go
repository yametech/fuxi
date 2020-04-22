package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
