package main

import "github.com/gin-gonic/gin"

func EventList(g *gin.Context) { workloadsAPI.ListEvent(g) }

func EventGet(g *gin.Context) { workloadsAPI.GetEvent(g) }
