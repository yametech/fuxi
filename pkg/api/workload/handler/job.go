package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/batch/v1"
	"net/http"
)

// Get Job
func (w *WorkloadsAPI) GetJob(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.job.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Job
func (w *WorkloadsAPI) ListJob(g *gin.Context) {
	list, err := w.job.List("", "", 0, 10000, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	jobList := &v1.JobList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, jobList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, jobList)
}
