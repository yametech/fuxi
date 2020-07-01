package main

import "github.com/gin-gonic/gin"

// PipelineRun doc
// @Summary tekton pipelineRun list
// @Description workload service for list pipelinRun
// @Tags PipelineRun
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/tekton.dev/v1alpha1/pipelineruns [get]
func PipelineRunList(g *gin.Context) { workloadsAPI.ListPipelineRun(g) }

// PipelineRun doc
// @Summary workload pipelineRun get
// @Description workload service for get a pipelinRun detail
// @Tags PipelineRun
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/tekton.dev/v1alpha1/namespaces/:namespace/pipelineruns/:name [get]
func PipelineRunGet(g *gin.Context) { workloadsAPI.GetPipelineRun(g) }

func PipelineRunCreate(g *gin.Context) { workloadsAPI.CreatePipelineRun(g) }
func PipelineRunUpdate(g *gin.Context) { workloadsAPI.UpdatePipelineRun(g) }
