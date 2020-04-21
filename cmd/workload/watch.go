package main

import "github.com/gin-gonic/gin"

// do not participate in swagger
func WatchStream(g *gin.Context) { workloadsAPI.WatchStream(g) }
