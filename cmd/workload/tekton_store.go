package main

import "github.com/gin-gonic/gin"

// TektonStore doc
// @Summary workload TektonStore list
// @Description workload service for list TektonStore
// @Tags Page
// @Accept mpfdx
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/tektonstores [get]
func TektonStoreList(g *gin.Context) { workloadsAPI.ListTektonStore(g) }

// TektonStore doc
// @Summary workload TektonStore get
// @Description workload service for get a TektonStore detail
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/tektonstores/{name} [get]
func TektonStoreGet(g *gin.Context) { workloadsAPI.GetTektonStore(g) }

// TektonStore doc
// @Summary workload TektonStore list
// @Description workload service for TektonStore
// @Tags Page
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/fuxi.nip.io/v1/namespaces/{namespace}/tektonstores/ [post]
func TektonStoreCreate(g *gin.Context) { workloadsAPI.CreateTektonStore(g) }
