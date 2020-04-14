package main

import "github.com/gin-gonic/gin"

// HorizontalPodAutoscaler doc
// @Summary workload horizontal pod autoscaler list
// @Description workload service for list horizontal pod autoscaler
// @Tags HorizontalPodAutoscaler
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/autoscaling/v2beta1/horizontalpodautoscalers [get]
func HorizontalPodAutoscalerList(g *gin.Context) { workloadsAPI.ListHorizontalPodAutoscaler(g) }

// HorizontalPodAutoscaler doc
// @Summary workload horizontal pod autoscaler get
// @Description workload service for get a horizontal pod autoscaler detail
// @Tags HorizontalPodAutoscaler
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name} [get]
func HorizontalPodAutoscalerGet(g *gin.Context) { workloadsAPI.GetHorizontalPodAutoscaler(g) }
