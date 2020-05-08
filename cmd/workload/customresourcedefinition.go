package main

import "github.com/gin-gonic/gin"

// GeneralCustomResourceDefinition doc
// @Summary workload general custom resource definition list
// @Description workload service for list general custom resource definition
// @Tags GeneralCustomResourceDefinition
// @Accept mpfd
// @Produce json
// @Param group query string true "group"
// @Param version query string true "version"
// @Param resource query string true "resource"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/crd/:group/:version/:resource [get]
func GeneralCustomResourceDefinitionList(g *gin.Context) {
	workloadsAPI.ListGeneralCustomResourceDefinition(g)
}

// CustomResourceDefinition doc
// @Summary workload custom resource definition list
// @Description workload service for list custom resource definition
// @Tags CustomResourceDefinition
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions [get]
func CustomResourceDefinitionList(g *gin.Context) { workloadsAPI.ListCustomResourceDefinition(g) }
