package main

import "github.com/gin-gonic/gin"

// BaseRoleUser doc
// @Summary base role user list
// @Description base service for list base role user
// @Tags BaseRoleUser
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/baseroleusers [get]
func BaseRoleUserList(g *gin.Context) { baseAPI.ListBaseRoleUser(g) }

// BaseRoleUser doc
// @Summary base role user get
// @Description base service for get a base roleuser detail
// @Tags BaseRoleUser
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseroleusers/{name} [get]
func BaseRoleUserGet(g *gin.Context) { baseAPI.GetBaseRoleUser(g) }

// BaseRoleUser doc
// @Summary base role user list
// @Description base service for base role user
// @Tags BaseRoleUser
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseroleusers/{name} [post]
func BaseRoleUserCreate(g *gin.Context) { baseAPI.CreateBaseRoleUser(g) }
