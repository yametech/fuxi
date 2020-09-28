package main

import "github.com/gin-gonic/gin"

// VirtualService doc
// @Summary  virtualservice list
// @Description workload service for list VirtualService
// @Tags VirtualService
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.istio.io/v1beta1/virtualservices [get]
func VirtualServiceList(g *gin.Context) { workloadsAPI.ListVirtulService(g) }

// VirtualService doc
// @Summary workload  get VirtualService
// @Description workload service for get a VirtualService detail
// @Tags VirtualService
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/networking.istio.io/v1beta1/namespaces/:namespace/virtualservices/:name [get]
func GetVirtualService(g *gin.Context) { workloadsAPI.GetVirtulService(g) }
