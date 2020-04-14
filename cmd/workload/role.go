package main

import "github.com/gin-gonic/gin"

// Role doc
// @Summary workload role list
// @Description workload service for list role
// @Tags Role
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/rbac.authorization.k8s.io/v1/roles [get]
func RoleList(g *gin.Context) { workloadsAPI.ListRole(g) }

// Role doc
// @Summary workload role get
// @Description workload service for get a role detail
// @Tags Role
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles/{name} [get]
func RoleGet(g *gin.Context) { workloadsAPI.GetRole(g) }
