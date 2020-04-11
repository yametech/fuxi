package main

import "github.com/gin-gonic/gin"

func ReplicaSetList(g *gin.Context) { workloadsAPI.ListReplicaset(g) }

func ReplicaSetGet(g *gin.Context) { workloadsAPI.GetReplicaset(g) }
