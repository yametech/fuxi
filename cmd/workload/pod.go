package main

import "github.com/gin-gonic/gin"

// Pod doc
// @Summary workload pod list
// @Description workload service for list pod
// @Tags Pod
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/pods [get]
func PodList(g *gin.Context) { workloadsAPI.ListPod(g) }

// Pod doc
// @Summary workload pod get
// @Description workload service for get a pod detail
// @Tags Pod
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/{namespace}/pods/{name} [get]
func PodGet(g *gin.Context) { workloadsAPI.GetPod(g) }

func PodAttach(g *gin.Context) { workloadsAPI.AttachPod(g) }
