package main

import "github.com/gin-gonic/gin"

// Gateway doc
// @Summary tekton gateway list
// @Description workload service for list gateway
// @Tags gateway
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.istio.io/v1beta1/gateways [get]
func GatewayList(g *gin.Context) { workloadsAPI.ListGateway(g) }

// Gateway doc
// @Summary workload  get GateWay
// @Description workload service for get a Gateway detail
// @Tags Gateway
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.istio.io/v1beta1/namespaces/:namespace/gateways/:name [get]
func GetGateway(g *gin.Context) { workloadsAPI.GetGateway(g) }
