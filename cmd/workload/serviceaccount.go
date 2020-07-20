package main

import "github.com/gin-gonic/gin"

// ServiceAccount doc
// @Summary workload service account list
// @Description workload service for list service account
// @Tags ServiceAccount
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/serviceaccounts [get]
func ServiceAccountList(g *gin.Context) { workloadsAPI.ListServiceAccount(g) }

// ServiceAccount doc
// @Summary workload service get
// @Description workload service for get a service account detail
// @Tags ServiceAccount
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/:namespace/serviceaccounts/:name [get]
func ServiceAccountGet(g *gin.Context) { workloadsAPI.GetServiceAccount(g) }

// ServiceAccount doc
// @Summary workload service get
// @Description workload service for get a service account detail
// @Tags ServiceAccount
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /serviceaccount/patch/:method [post]
func ServiceAccountPatchSecret(g *gin.Context) { workloadsAPI.PatchSecretServiceAccount(g) }
