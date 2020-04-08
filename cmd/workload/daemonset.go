package main

import "github.com/gin-gonic/gin"

// DaemonSet doc
// @Summary workload daemonSet list
// @Description workload service for list cronJob
// @Tags DaemonSet
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/batch/v1beta1/daemonset [get]
func DaemonSetList(g *gin.Context) { workloadsAPI.ListDaemonSet(g) }

// DaemonSet doc
// @Summary workload daemonSet get
// @Description workload service for get a cronJob detail
// @Tags DaemonSet
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/v1/:ns/daemonset/:name/get [get]
func DaemonSetGet(g *gin.Context) { workloadsAPI.GetDaemonSet(g) }
