package main

import "github.com/gin-gonic/gin"

// ClusterRoleBind doc
// @Summary workload configMaps list
// @Description workload service for list ClusterRoleBind
// @Tags ClusterRoleBind
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/rbac.authorization.k8s.io/v1/clusterrolebindings [get]
func ClusterRoleBindList(g *gin.Context) { workloadsAPI.ListClusterRoleBind(g) }

// ClusterRoleBind doc
// @Summary workload configMaps get
// @Description workload service for get a ClusterRoleBind detail
// @Tags ClusterRoleBind
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/clusterrolebindings/:name [get]
func ClusterRoleBindGet(g *gin.Context) { workloadsAPI.GetClusterRoleBind(g) }
