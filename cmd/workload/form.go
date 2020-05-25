package main

import "github.com/gin-gonic/gin"

// Form doc
// @Summary workload form list
// @Description workload service for list form
// @Tags Form
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/forms [get]
func FormList(g *gin.Context) { workloadsAPI.ListForm(g) }

// Form doc
// @Summary workload form get
// @Description workload service for get a form detail
// @Tags Form
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/forms/{name} [get]
func FormGet(g *gin.Context) { workloadsAPI.GetForm(g) }

// Form doc
// @Summary workload form list
// @Description workload service for form
// @Tags Form
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/forms/{name} [post]
func FormCreate(g *gin.Context) { workloadsAPI.CreateForm(g) }
