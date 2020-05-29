package main

import "github.com/gin-gonic/gin"

// Task doc
// @Summary tekton task list
// @Description workload service for list task
// @Tags Task
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/tekton.dev/v1alpha1/tasks [get]
func TaskList(g *gin.Context) {
	workloadsAPI.ListTask(g)
}

// Task doc
// @Summary workload Task get
// @Description workload service for get a task detail
// @Tags Task
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/tekton.dev/v1alpha1/namespaces/:namespace/tasks/:name [get]
func TaskGet(g *gin.Context) {
	workloadsAPI.GetTask(g)
}
