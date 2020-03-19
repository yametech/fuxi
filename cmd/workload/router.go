package main

import (
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

type Pod corev1.PodTemplate

// Pod doc
// @Summary workload pod list
// @Description workload service for list pod
// @Tags Pod
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/v1/:ns/pod/list [get]
func (p *Pod) List(g *gin.Context) {

}

// Pod doc
// @Summary workload pod get
// @Description workload service for get a pod detail
// @Tags Pod
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /workload/v1/:ns/pod/:name/get [get]
func (p *Pod) Get(g *gin.Context) {

}

func (p *Pod) Update(g *gin.Context) {

}

func (p *Pod) Delete(g *gin.Context) {

}

func (p *Pod) Attach(g *gin.Context) {

}

func (p *Pod) Log(g *gin.Context) {
}
