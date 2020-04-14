package main

import "github.com/gin-gonic/gin"

// PersistentVolumeClaims doc
// @Summary workload persistent volume claims list
// @Description workload service for list network policy
// @Tags PersistentVolumeClaims
// @Accept mpfd
// @Produce json
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/persistentvolumeclaims [get]
func PersistentVolumeClaimsList(g *gin.Context) { workloadsAPI.ListPersistentVolumeClaims(g) }

// PersistentVolumeClaims doc
// @Summary workload persistent volume claims get
// @Description workload service for get a persistent volume claims detail
// @Tags PersistentVolumeClaims
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/namespaces/{namespace}/persistentvolumeclaims/{name} [get]
func PersistentVolumeClaimsGet(g *gin.Context) { workloadsAPI.GetPersistentVolumeClaims(g) }
