package main

import "github.com/gin-gonic/gin"

// TektonGraph doc
// @Summary workload tektongraph list
// @Description workload service for list tektongraph
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/tektongraphs [get]
func TektonGraphList(g *gin.Context) { workloadsAPI.ListTektonGraph(g) }

// TektonGraph doc
// @Summary workload tektongraph get
// @Description workload service for get a tektongraph detail
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/tektongraphs/{name} [get]
func TektonGraphGet(g *gin.Context) { workloadsAPI.GetTektonGraph(g) }

// TektonGraph doc
// @Summary workload tektongraph list
// @Description workload service for tektongraph
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/tektongraphs/ [post]
func TektonGraphCreate(g *gin.Context) { workloadsAPI.CreateTektonGraph(g) }
