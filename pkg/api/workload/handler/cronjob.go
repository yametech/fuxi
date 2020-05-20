package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/api/batch/v1beta1"
)

// Get CronJob
func (w *WorkloadsAPI) GetCronJob(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.cronJob.Get(namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List CronJob
func (w *WorkloadsAPI) ListCronJob(g *gin.Context) {
	list, err := w.cronJob.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	cronJobList := &v1beta1.CronJobList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	if err = json.Unmarshal(marshalData, cronJobList); err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, cronJobList)
}
