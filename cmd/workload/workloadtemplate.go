package main

import "github.com/gin-gonic/gin"

// Workloads doc
// @Summary workloads template list
// @Description workload service for list workloads template
// @Tags Workloads
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/workloads [get]
func WorkloadsTemplateList(g *gin.Context) { workloadsAPI.ListWorkloadsTemplate(g) }

// Workloads doc
// @Summary workloads template get
// @Description workload service for get a workloads template detail
// @Tags Workloads
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/workloads/namespaces/{namespace}/{name} [get]
func WorkloadsTemplateGet(g *gin.Context) { workloadsAPI.GetWorkloadsTemplate(g) }

func WorkloadsTemplateListSharedNamespace(g *gin.Context) {
	workloadsAPI.ListShareNamespacedWorkloadsTemplate(g)
}

func WorkloadsTemplateCreate(g *gin.Context) { workloadsAPI.PostWorkloadsTemplate(g) }
