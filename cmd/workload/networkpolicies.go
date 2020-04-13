package main

import "github.com/gin-gonic/gin"

// NetworkPolicy doc
// @Summary workload network policy list
// @Description workload service for list network policy
// @Tags NetworkPolicy
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.k8s.io/v1/networkpolicies [get]
func NetworkPolicyList(g *gin.Context) { workloadsAPI.ListNetworkPolicy(g) }

// NetworkPolicy doc
// @Summary workload network policy get
// @Description workload service for get a network policy detail
// @Tags NetworkPolicy
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.k8s.io/v1/namespaces/{namespace}/networkpolicies/{name} [get]
func NetworkPolicyGet(g *gin.Context) { workloadsAPI.GetNetworkPolicy(g) }
