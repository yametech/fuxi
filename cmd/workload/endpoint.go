package main

import "github.com/gin-gonic/gin"

// Endpoint doc
// @Summary workload form render list
// @Description workload service for list form render
// @Tags Endpoint
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/endpoints [get]
func EndpointList(g *gin.Context) { workloadsAPI.ListEndpoint(g) }

// Endpoint doc
// @Summary workload form render get
// @Description workload service for get a form render detail
// @Tags Endpoint
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/{namespace}/endpoints/{name} [get]
func EndpointGet(g *gin.Context) { workloadsAPI.GetEndpoint(g) }
