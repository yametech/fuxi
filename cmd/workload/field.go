package main

import "github.com/gin-gonic/gin"

// Field doc
// @Summary workload field list
// @Description workload service for list field
// @Tags Field
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/fields [get]
func FieldList(g *gin.Context) { workloadsAPI.ListField(g) }

// Field doc
// @Summary workload field get
// @Description workload service for get a field detail
// @Tags Field
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/fields/{name} [get]
func FieldGet(g *gin.Context) { workloadsAPI.GetField(g) }

// Field doc
// @Summary workload field list
// @Description workload service for field
// @Tags Field
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/fields/{name} [post]
func FieldCreate(g *gin.Context) { workloadsAPI.CreateField(g) }
