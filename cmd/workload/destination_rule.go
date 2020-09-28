package main

import "github.com/gin-gonic/gin"

// DestinationRule doc
// @Summary  DestinationRule list
// @Description workload service for list DestinationRule
// @Tags DestinationRule
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.istio.io/v1beta1/destinationrules [get]
func DestinationRuleList(g *gin.Context) { workloadsAPI.ListDestinationRule(g) }

// DestinationRule doc
// @Summary workload  get DestinationRule
// @Description workload service for get a DestinationRule detail
// @Tags DestinationRule
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/networking.istio.io/v1beta1/namespaces/:namespace/destinationrules/:name [get]
func GetDestinationRule(g *gin.Context) { workloadsAPI.GetDestinationRule(g) }
