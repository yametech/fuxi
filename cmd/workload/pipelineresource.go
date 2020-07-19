package main

import "github.com/gin-gonic/gin"

// PipelineResource doc
// @Summary tekton PipelineResource list
// @Description workload service for list PipelineResource
// @Tags PipelineResource
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/tekton.dev/v1alpha1/pipelineresources [get]
func PipelineResourceList(g *gin.Context) { workloadsAPI.ListPipelineResource(g) }

// PipelineResource doc
// @Summary workload PipelineResource get
// @Description workload service for get a PipelineResource detail
// @Tags PipelineResource
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineresources/:name [get]
func PipelineResourceGet(g *gin.Context) { workloadsAPI.GetPipelineResource(g) }
