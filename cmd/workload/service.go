package main

import "github.com/gin-gonic/gin"

// Service doc
// @Summary workload service list
// @Description workload service for list service
// @Tags Service
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/services [get]
func ServiceList(g *gin.Context) { workloadsAPI.ListService(g) }

// Service doc
// @Summary workload service get
// @Description workload service for get a service detail
// @Tags Service
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/:namespace/services/:name [get]
func ServiceGet(g *gin.Context) { workloadsAPI.GetService(g) }
