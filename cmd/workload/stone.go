package main

import "github.com/gin-gonic/gin"

// stones doc
// @Summary workload stones list
// @Description workload service for list stones
// @Tags stones
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/nuwa.nip.io/v1/stones [get]
func StoneList(g *gin.Context) { workloadsAPI.ListStone(g) }

// stones doc
// @Summary workload stones get
// @Description workload service for get a stones detail
// @Tags stones
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/nuwa.nip.io/v1/namespaces/:namespace/stones/:name [get]
func StoneGet(g *gin.Context) { workloadsAPI.GetStone(g) }
