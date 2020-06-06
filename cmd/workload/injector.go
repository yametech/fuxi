package main

import "github.com/gin-gonic/gin"

// Injectors doc
// @Summary workload Injectors list
// @Description workload service for list Injectors
// @Tags Injectors
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/nuwa.nip.io/v1/Injectors [get]
func InjectorList(g *gin.Context) { workloadsAPI.ListInjector(g) }

// Injectors doc
// @Summary workload Injectors get
// @Description workload service for get a Injectors detail
// @Tags Injectors
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/nuwa.nip.io/v1/namespaces/:namespace/Injectors/:name [get]
func InjectorGet(g *gin.Context) { workloadsAPI.GetInjector(g) }
