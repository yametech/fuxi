package main

import "github.com/gin-gonic/gin"

// RoleBinding doc
// @Summary workload role binding list
// @Description workload service for list role binding
// @Tags RoleBinding
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/rbac.authorization.k8s.io/v1/rolebindings [get]
func RoleBindingList(g *gin.Context) { workloadsAPI.ListRoleBinding(g) }

// RoleBinding doc
// @Summary workload role binding get
// @Description workload service for get a role binding detail
// @Tags RoleBinding
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings/{name} [get]
func RoleBindingGet(g *gin.Context) { workloadsAPI.GetRoleBinding(g) }
