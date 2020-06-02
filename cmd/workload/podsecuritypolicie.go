package main

import "github.com/gin-gonic/gin"

// PodSecurityPolicie doc
// @Summary workload PodSecurityPolicie list
// @Description workload service for list podsecuritypolicies
// @Tags PodSecurityPolicie
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/policy/v1beta1/podsecuritypolicies [get]
func PodSecurityPolicieList(g *gin.Context) { workloadsAPI.ListPodSecurityPolicie(g) }

// PodSecurityPolicie doc
// @Summary workload podsecuritypolicies get
// @Description workload service for get a podsecuritypolicies detail
// @Tags PodSecurityPolicie
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/policy/v1beta1/namespaces/{namespace}/podsecuritypolicies/{name} [get]
func PodSecurityPolicieGet(g *gin.Context) { workloadsAPI.GetPodSecurityPolicie(g) }
