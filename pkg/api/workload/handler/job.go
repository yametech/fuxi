package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/batch/v1"
	"net/http"
)

// Get Job
func (w *WorkloadsAPI) GetJob(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.job.Get(dyn.ResourceJob, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Job
func (w *WorkloadsAPI) ListJob(g *gin.Context) {
	list, _ := w.job.List(dyn.ResourceJob, "", "", 0, 10000, nil)
	jobList := &v1.JobList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, jobList)
	g.JSON(http.StatusOK, jobList)
}