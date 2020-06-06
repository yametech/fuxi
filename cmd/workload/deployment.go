package main

import "github.com/gin-gonic/gin"

// Deployment doc
// @Summary workload deployment list
// @Description workload service for list cronJob
// @Tags Deployment
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/batch/v1beta1/deployment [get]
func DeploymentList(g *gin.Context) { workloadsAPI.ListDeployment(g) }

// Deployment doc
// @Summary workload cronJob get
// @Description workload service for get a cronJob detail
// @Tags Deployment
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/v1/:ns/deployment/:name/get [get]
func DeploymentGet(g *gin.Context) { workloadsAPI.GetDeployment(g) }

