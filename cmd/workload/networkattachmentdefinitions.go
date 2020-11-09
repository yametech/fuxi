package main

import "github.com/gin-gonic/gin"

// NetworkAttachmentDefinition doc
// @Summary workload network-attachment-definitions list
// @Description workload service for list network-attachment-definitions
// @Tags NetworkAttachmentDefinition
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/k8s.cni.cncf.io/v1/network-attachment-definitions [get]
func NetworkAttachmentDefinitionList(g *gin.Context) { workloadsAPI.ListNetworkAttachmentDefinition(g) }

// NetworkAttachmentDefinition doc
// @Summary workload network-attachment-definitions get
// @Description workload service for get a network-attachment-definitions detail
// @Tags NetworkAttachmentDefinition
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/k8s.cni.cncf.io/v1/namespaces/{namespace}/network-attachment-definitions/{name} [get]
func NetworkAttachmentDefinitionGet(g *gin.Context) { workloadsAPI.GetNetworkAttachmentDefinition(g) }
