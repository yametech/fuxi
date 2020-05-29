package main

import "github.com/gin-gonic/gin"

// PersistentVolume doc
// @Summary workload persistent volume list
// @Description workload service for list persistent volume
// @Tags PersistentVolume
// @Accept mpfd
// @Produce json
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/persistentvolumes [get]
func PersistentVolumeList(g *gin.Context) { workloadsAPI.ListPersistentVolume(g) }

// PersistentVolume doc
// @Summary workload persistent volume get
// @Description workload service for get a persistent volume detail
// @Tags PersistentVolume
// @Accept mpfd
// @Produce json
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/persistentvolumes/{name} [get]
func PersistentVolumeGet(g *gin.Context) { workloadsAPI.GetPersistentVolume(g) }

// PersistentVolume doc
// @Summary workload persistent volume get
// @Description workload service for get a persistent volume detail
// @Tags PersistentVolume
// @Accept mpfd
// @Produce json
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/api/v1/persistentvolumes/{name} [delete]
func PersistentVolumeDelete(g *gin.Context) { workloadsAPI.DeletePersistentVolume(g) }
