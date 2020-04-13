package main

import "github.com/gin-gonic/gin"

// ResourceQuota doc
// @Summary workload resource quota list
// @Description workload service for list resource quota
// @Tags ResourceQuota
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/resourcequotas [get]
func ResourceQuotaList(g *gin.Context) { workloadsAPI.ListResourceQuota(g) }

// ResourceQuota doc
// @Summary workload resource quota get
// @Description workload service for get a resource quota detail
// @Tags ResourceQuota
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/:namespace/resourcequotas/:name [get]
func ResourceQuotaGet(g *gin.Context) { workloadsAPI.GetResourceQuota(g) }
