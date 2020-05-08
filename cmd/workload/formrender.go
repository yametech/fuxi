package main

import "github.com/gin-gonic/gin"

// FormRender doc
// @Summary workload form render list
// @Description workload service for list form render
// @Tags FormRender
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/formrenders [get]
func FormRenderList(g *gin.Context) { workloadsAPI.ListFormRender(g) }

// FormRender doc
// @Summary workload form render get
// @Description workload service for get a form render detail
// @Tags FormRender
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/formrender/{name} [get]
func FormRenderGet(g *gin.Context) { workloadsAPI.GetFormRender(g) }
