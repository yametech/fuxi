package main

import "github.com/gin-gonic/gin"

func ReplicaSetList(g *gin.Context) { workloadsAPI.ListReplicaSet(g) }

func ReplicaSetGet(g *gin.Context) { workloadsAPI.GetReplicaSet(g) }
