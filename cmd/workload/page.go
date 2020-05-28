package main

import "github.com/gin-gonic/gin"

// Page doc
// @Summary workload page list
// @Description workload service for list page
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/pages [get]
func PageList(g *gin.Context) { workloadsAPI.ListPage(g) }

// Page doc
// @Summary workload page get
// @Description workload service for get a page detail
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/pages/{name} [get]
func PageGet(g *gin.Context) { workloadsAPI.GetPage(g) }

// Page doc
// @Summary workload page list
// @Description workload service for page
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/pages/{name} [post]
func PageCreate(g *gin.Context) { workloadsAPI.CreatePage(g) }
