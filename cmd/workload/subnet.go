package main

import (
	"github.com/gin-gonic/gin"
)

// SubNet doc
// @Summary workload subnet list
// @Description workload service for list subnet
// @Tags SubNet
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/kubeovn.io/v1/subnets [get]
func SubNetList(g *gin.Context) { workloadsAPI.ListSubNet(g) }

// SubNet doc
// @Summary workload subnet get
// @Description workload service for get a subnet detail
// @Tags SubNet
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/kubeovn.io/v1/namespaces/{namespace}/subnets/{name} [get]
func SubNetGet(g *gin.Context) { workloadsAPI.GetSubNet(g) }

// SubNet doc
// @Summary workload subnet list
// @Description workload service for subnet
// @Tags SubNet
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/kubeovn.io/v1/subnets/ [post]
func SubNetCreate(g *gin.Context) { workloadsAPI.CreateSubNet(g) }
