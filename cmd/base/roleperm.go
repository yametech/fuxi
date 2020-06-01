package main

import "github.com/gin-gonic/gin"

// BaseRolePerm doc
// @Summary base role perm list
// @Description base service for list base role perm
// @Tags BaseRolePerm
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/baseroleperms [get]
func BaseRolePermList(g *gin.Context) { baseAPI.ListBaseRolePerm(g) }

// BaseRolePerm doc
// @Summary base role perm get
// @Description base service for get a base roleperm detail
// @Tags BaseRolePerm
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseroleperms/{name} [get]
func BaseRolePermGet(g *gin.Context) { baseAPI.GetBaseRolePerm(g) }

// BaseRolePerm doc
// @Summary base role perm list
// @Description base service for base role perm
// @Tags BaseRolePerm
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseroleperms/{name} [post]
func BaseRolePermCreate(g *gin.Context) { baseAPI.CreateBaseRolePerm(g) }
