package main

import "github.com/gin-gonic/gin"

// BaseUser doc
// @Summary base role list
// @Description base service for list base role
// @Tags BaseUser
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/baseusers [get]
func BaseUserList(g *gin.Context) { baseAPI.ListBaseUser(g) }

// BaseUser doc
// @Summary base role get
// @Description base service for get a base role detail
// @Tags BaseUser
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseusers/{name} [get]
func BaseUserGet(g *gin.Context) { baseAPI.GetBaseUser(g) }

// BaseUser doc
// @Summary base role list
// @Description base service for base role
// @Tags BaseUser
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseusers/{name} [post]
func BaseUserCreate(g *gin.Context) { baseAPI.CreateBaseUser(g) }
