package main

import "github.com/gin-gonic/gin"

// Namespaces doc
// @Summary workload namespaces list
// @Description workload service for list network policy
// @Tags Namespaces
// @Accept mpfd
// @Produce json
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces [get]
func NamespaceList(g *gin.Context) { workloadsAPI.ListNamespace(g) }

// Namespaces doc
// @Summary workload namespaces list
// @Description workload service for list network policy
// @Tags Namespaces
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/:namespace [get]
func NamespaceGet(g *gin.Context) { workloadsAPI.GetNamespace(g) }

// Namespaces doc
// @Summary workload namespaces list
// @Description workload service for list network policy
// @Tags Namespaces
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/:namespace [delete]
func NamespaceDelete(g *gin.Context) { workloadsAPI.DeleteNamespace(g) }

// Namespaces doc
// @Summary workload namespaces list
// @Description workload service for list network policy
// @Tags Namespaces
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/:namespace [post]
func NamespaceCreate(g *gin.Context) { workloadsAPI.CreateNamespace(g) }
