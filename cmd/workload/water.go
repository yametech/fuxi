package main

import "github.com/gin-gonic/gin"

// Waters doc
// @Summary workload Waters list
// @Description workload service for list Waters
// @Tags Waters
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/nuwa.nip.io/v1/Waters [get]
func WaterList(g *gin.Context) { workloadsAPI.ListWater(g) }

// Waters doc
// @Summary workload Waters get
// @Description workload service for get a Waters detail
// @Tags Waters
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/nuwa.nip.io/v1/namespaces/:namespace/Waters/:name [get]
func WaterGet(g *gin.Context) { workloadsAPI.GetWater(g) }
