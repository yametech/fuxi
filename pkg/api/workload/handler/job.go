package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/batch/v1"
)

// Get Job
func (w *WorkloadsAPI) GetJob(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.job.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Job
func (w *WorkloadsAPI) ListJob(g *gin.Context) {
	list, err := resourceList(g, w.job)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	jobList := &v1.JobList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, jobList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, jobList)
}
