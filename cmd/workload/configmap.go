package main

import "github.com/gin-gonic/gin"

// ConfigMaps doc
// @Summary workload configMaps list
// @Description workload service for list configMaps
// @Tags ConfigMaps
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/configmaps [get]
func ConfigMapsList(g *gin.Context) { workloadsAPI.ListConfigMaps(g) }

// ConfigMap doc
// @Summary workload configMaps get
// @Description workload service for get a configMaps detail
// @Tags ConfigMap
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/:namespace/configmaps/:name [get]
func ConfigMapsGet(g *gin.Context) { workloadsAPI.GetConfigMaps(g) }
