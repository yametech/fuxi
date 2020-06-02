package main

import "github.com/gin-gonic/gin"

// Permission doc
// @Summary base permission
// @Description base service for list base permission
// @Tags Permission
// @Accept mpfd
// @Produce json
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/permission_transfer/{value} [get]
func BasePermissionTransferList(g *gin.Context) { baseAPI.PermissionTransfer(g) }
