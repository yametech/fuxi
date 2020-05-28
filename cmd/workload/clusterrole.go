package main

import "github.com/gin-gonic/gin"

// ClusterRole doc
// @Summary workload role list
// @Description workload service for list role
// @Tags ClusterRole
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/rbac.authorization.k8s.io/v1/clusterroles [get]
func ClusterRoleList(g *gin.Context) { workloadsAPI.ListClusterRole(g) }

// ClusterRole doc
// @Summary workload ClusterRole get
// @Description workload service for get a ClusterRole detail
// @Tags ClusterRole
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/clusterroles/:name [get]
func ClusterRoleGet(g *gin.Context) { workloadsAPI.GetClusterRole(g) }
