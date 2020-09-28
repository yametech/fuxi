package main

import "github.com/gin-gonic/gin"

// WorkloadEntry doc
// @Summary  WorkloadEntry list
// @Description workload service for list WorkloadEntry
// @Tags WorkloadEntry
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.istio.io/v1beta1/workloadentrys [get]
func WorkloadEntryList(g *gin.Context) { workloadsAPI.GetWorkloadEntry(g) }

// WorkloadEntry doc
// @Summary workload  get WorkloadEntry
// @Description workload service for get a WorkloadEntry detail
// @Tags WorkloadEntry
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/networking.istio.io/v1beta1/namespaces/:namespace/workloadentrys/:name [get]
func GetWorkloadEntry(g *gin.Context) { workloadsAPI.GetWorkloadEntry(g) }
