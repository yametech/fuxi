package main

import "github.com/gin-gonic/gin"

// GiteaWebHook doc
// @Summary gitea webhook
// @Description workload service for gitea webhook
// @Tags Form
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /webhook/gitea/ [post]
func TriggerGiteaWebHook(g *gin.Context) { workloadsAPI.TriggerGiteaWebHook(g) }
