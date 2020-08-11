package main

import "github.com/gin-gonic/gin"

// TektonWebHook doc
// @Summary workload tektongraph list
// @Description workload service for list tektongraph
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/tektonwebhooks [get]
func TektonWebHookList(g *gin.Context) { workloadsAPI.ListTektonWebHook(g) }

// TektonWebHook doc
// @Summary workload tektongraph get
// @Description workload service for get a tektongraph detail
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/tektonwebhooks/{name} [get]
func TektonWebHookGet(g *gin.Context) { workloadsAPI.GetTektonWebHook(g) }

// TektonWebHook doc
// @Summary workload tektongraph list
// @Description workload service for tektongraph
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/tektonwebhooks/ [post]
func TektonWebHookCreate(g *gin.Context) { workloadsAPI.CreateTektonWebHook(g) }
