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
// @Router /base/permission_list [get]
func BasePermissionList(g *gin.Context) { baseAPI.PermissionList(g) }

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

// Permission doc
// @Summary base permission
// @Description base service for list base permission
// @Tags Permission
// @Accept mpfd
// @Produce json
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/permission_auth_value [post]
func BasePermissionAuthorizeValue(g *gin.Context) { baseAPI.PermissionAuthorizeValue(g) }
