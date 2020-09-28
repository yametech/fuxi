package main

import "github.com/gin-gonic/gin"

// Sidecar doc
// @Summary  Sidecar list
// @Description workload service for list Sidecar
// @Tags Sidecar
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/networking.istio.io/v1beta1/sidecars [get]
func SidecarList(g *gin.Context) { workloadsAPI.ListSidecar(g) }

// Sidecar doc
// @Summary workload  get Sidecar
// @Description workload service for get a Sidecar detail
// @Tags Sidecar
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/networking.istio.io/v1beta1/namespaces/:namespace/sidecar/:name [get]
func GetSidecar(g *gin.Context) { workloadsAPI.GetSidecar(g) }
