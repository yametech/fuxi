package main

import "github.com/gin-gonic/gin"

// StatefulSet doc
// @Summary workload statefulSet list
// @Description workload service for list cronJob
// @Tags StatefulSet
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/apps/v1/statefulsets [get]
func StatefulSetList(g *gin.Context) { workloadsAPI.ListStatefulSet(g) }

// StatefulSet doc
// @Summary workload statefulSet get
// @Description workload service for get a cronJob detail
// @Tags StatefulSet
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/apps/v1/namespaces/:namespace/statefulsets/:name [get]
func StatefulSetGet(g *gin.Context) { workloadsAPI.GetStatefulSet(g) }
