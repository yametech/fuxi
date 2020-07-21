package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/pkg/api/common"
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
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List CronJob
func (w *WorkloadsAPI) ListCronJob(g *gin.Context) {
	list, err := resourceList(g, w.cronJob)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	cronJobList := &v1beta1.CronJobList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	if err = json.Unmarshal(marshalData, cronJobList); err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, cronJobList)
}
