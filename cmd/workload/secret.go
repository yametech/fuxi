package main

import "github.com/gin-gonic/gin"

// Secret doc
// @Summary workload secret list
// @Description workload service for list secret
// @Tags Secret
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/secrets [get]
func SecretList(g *gin.Context) { workloadsAPI.ListSecret(g) }

// Secret doc
// @Summary workload secret list
// @Description workload service for list secret
// @Tags Secret
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/ops-secrets [get]
func OpsSecretList(g *gin.Context) { workloadsAPI.ListOpsSecret(g) }

// Secret doc
// @Summary workload secret get
// @Description workload service for get a secret detail
// @Tags ConfigMap
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/{namespace}/secrets/{name} [get]
func SecretGet(g *gin.Context) { workloadsAPI.GetSecret(g) }

// Secret doc
// @Summary workload secret get
// @Description workload service for get a secret detail
// @Tags ConfigMap
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/{namespace}/secrets/ [post]
// @Router /workload/api/v1/namespaces/{namespace}/ops-secrets/ [post]
func SecretCreate(g *gin.Context) { workloadsAPI.CreateSecret(g) }

// Secret doc
// @Summary workload secret get
// @Description workload service for get a secret detail
// @Tags ConfigMap
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/{namespace}/secrets/ [put]
// @Router /workload/api/v1/namespaces/{namespace}/ops-secrets/ [put]
func SecretUpdate(g *gin.Context) { workloadsAPI.UpdateSecret(g) }
