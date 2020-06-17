package main

import "github.com/gin-gonic/gin"

// PipelineResourc doc
// @Summary tekton PipelineResourc list
// @Description workload service for list PipelineResourc
// @Tags PipelineResourc
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/tekton.dev/v1alpha1/pipelineresources [get]
func PipelineResourceList(g *gin.Context) { workloadsAPI.ListPipelineResource(g) }

// PipelineResourc doc
// @Summary workload PipelineResourc get
// @Description workload service for get a PipelineResourc detail
// @Tags PipelineResourc
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/tekton.dev/v1alpha1/namespaces/:namespace/pipelineresources/:name [get]
func PipelineResourceGet(g *gin.Context) { workloadsAPI.GetPipelineResource(g) }
