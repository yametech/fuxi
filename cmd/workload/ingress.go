package main

import "github.com/gin-gonic/gin"

// Ingress doc
// @Summary workload ingress list
// @Description workload service for list ingress
// @Tags Ingress
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/extensions/v1beta1/ingresses [get]
func IngressList(g *gin.Context) { workloadsAPI.ListIngress(g) }

// Ingress doc
// @Summary workload ingress get
// @Description workload service for get a ingress detail
// @Tags Ingress
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/extensions/v1beta1/namespaces/:namespace/ingresses/:name [get]
func IngressGet(g *gin.Context) { workloadsAPI.GetIngress(g) }

//func IngressCreate(g *gin.Context) { workloadsAPI.CreateIngress(g) }
