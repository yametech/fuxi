package main

import "github.com/gin-gonic/gin"

// CronJob doc
// @Summary workload cronJob list
// @Description workload service for list cronJob
// @Tags CronJob
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/batch/v1beta1/cronjobs [get]
func CronJobList(g *gin.Context) { workloadsAPI.ListCronJob(g) }

// CronJob doc
// @Summary workload cronJob get
// @Description workload service for get a cronJob detail
// @Tags CronJob
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/v1/:ns/cronjobs/:name/get [get]
func CronJobGet(g *gin.Context) { workloadsAPI.GetCronJob(g) }
