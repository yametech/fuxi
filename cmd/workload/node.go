package main

import "github.com/gin-gonic/gin"

func NodeList(g *gin.Context) { workloadsAPI.ListNode(g) }

func NodeGet(g *gin.Context) { workloadsAPI.GetNode(g) }
