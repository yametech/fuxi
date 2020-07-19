package main

import "github.com/gin-gonic/gin"

// Pipeline doc
// @Summary tekton pipeline list
// @Description workload service for list pipeline
// @Tags Pipeline
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/tekton.dev/v1alpha1/pipelines [get]
func PipelineList(g *gin.Context) { workloadsAPI.ListPipeline(g) }

// Pipeline doc
// @Summary workload Pipeline get
// @Description workload service for get a Pipeline detail
// @Tags Pipeline
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/tekton.dev/v1alpha1/namespaces/:namespace/pipelines/:name [get]
func PipelineGet(g *gin.Context) { workloadsAPI.GetPipeline(g) }
