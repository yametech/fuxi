package main

import "github.com/gin-gonic/gin"

// ServiceEntry doc
// @Summary  ServiceEntry list
// @Description workload service for list ServiceEntry
// @Tags ServiceEntry
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.istio.io/v1beta1/serviceentries [get]
func ServiceEntryList(g *gin.Context) { workloadsAPI.ListServiceEntry(g) }

// ServiceEntry doc
// @Summary workload  get ServiceEntry
// @Description workload service for get a ServiceEntry detail
// @Tags ServiceEntry
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/networking.istio.io/v1beta1/namespaces/:namespace/serviceentries/:name [get]
func GetServiceEntry(g *gin.Context) { workloadsAPI.GetServiceEntry(g) }
