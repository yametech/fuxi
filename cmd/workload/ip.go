package main

import "github.com/gin-gonic/gin"

// IP doc
// @Summary workload ip list
// @Description workload service for list ip
// @Tags IP
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/kubeovn.io/v1/ips [get]
func IPList(g *gin.Context) { workloadsAPI.ListIP(g) }

// IP doc
// @Summary workload ip get
// @Description workload service for get a ip detail
// @Tags IP
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/kubeovn.io/v1/namespaces/{namespace}/ips/{name} [get]
func IPGet(g *gin.Context) { workloadsAPI.GetIP(g) }
