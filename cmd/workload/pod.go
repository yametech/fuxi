package main

import "github.com/gin-gonic/gin"

func PodList(g *gin.Context) { workloadsAPI.ListPod(g) }

func PodGet(g *gin.Context) { workloadsAPI.GetPod(g) }

func PodAttach(g *gin.Context) { workloadsAPI.AttachPod(g) }
