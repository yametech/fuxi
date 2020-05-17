package main

import "github.com/gin-gonic/gin"

// StatefulSet1 doc
// @Summary workload statefulSet1 list
// @Description workload service for list cronJob
// @Tags StatefulSet1
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/nuwa.nip.io/v1/statefulsets [get]
func StatefulSet1List(g *gin.Context) { workloadsAPI.ListStatefulSet1(g) }

// StatefulSet1 doc
// @Summary workload statefulSet1 get
// @Description workload service for get a cronJob detail
// @Tags StatefulSet1
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/nuwa.nip.io/v1/namespaces/:namespace/statefulsets/:name [get]
func StatefulSet1Get(g *gin.Context) { workloadsAPI.GetStatefulSet1(g) }
