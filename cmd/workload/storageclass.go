package main

import "github.com/gin-gonic/gin"

// StorageClass doc
// @Summary workload storage class list
// @Description workload service for list storage class
// @Tags StorageClass
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/storage.k8s.io/v1beta1/storageclasses [get]
func StorageClassList(g *gin.Context) { workloadsAPI.ListStorageClass(g) }

// StorageClass doc
// @Summary workload storage class get
// @Description workload service for get a storage class detail
// @Tags StorageClass
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/storage.k8s.io/v1beta1/storageclasses/{name} [get]
func StorageClassGet(g *gin.Context) { workloadsAPI.GetStorageClass(g) }
