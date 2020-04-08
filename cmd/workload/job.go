package main

import "github.com/gin-gonic/gin"

// Job doc
// @Summary workload daemonSet list
// @Description workload service for list cronJob
// @Tags Job
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/batch/v1/jobs [get]
func JobList(g *gin.Context) { workloadsAPI.ListJob(g) }

// Job doc
// @Summary workload daemonSet get
// @Description workload service for get a cronJob detail
// @Tags Job
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/batch/v1/namespaces/:namespace/jobs/:name [get]
func JobGet(g *gin.Context) { workloadsAPI.GetJob(g) }
