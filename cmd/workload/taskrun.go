package main

import "github.com/gin-gonic/gin"

// TaskRun doc
// @Summary tekton taskrun list
// @Description workload service for list taskrun
// @Tags TaskRun
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/tekton.dev/v1alpha1/taskruns [get]
func TaskRunList(g *gin.Context) {
	workloadsAPI.ListTaskRun(g)
}

// TaskRun doc
// @Summary workload TaskRun get
// @Description workload service for get a taskrun detail
// @Tags TaskRun
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/tekton.dev/v1alpha1/namespaces/:namespace/taskruns/:name [get]
func TaskRunGet(g *gin.Context) {
	workloadsAPI.GetTaskRun(g)
}
