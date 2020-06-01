package main

import "github.com/gin-gonic/gin"

// BaseRole doc
// @Summary base role list
// @Description base service for list base role
// @Tags BaseRole
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/baseroles [get]
func BaseRoleList(g *gin.Context) { baseAPI.ListBaseRole(g) }

// BaseRole doc
// @Summary base role get
// @Description base service for get a base role detail
// @Tags BaseRole
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/baseroles/{name} [get]
func BaseRoleGet(g *gin.Context) { baseAPI.GetBaseRole(g) }

// BaseRole doc
// @Summary base role list
// @Description base service for base role
// @Tags BaseRole
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/baseroles/{name} [post]
func BaseRoleCreate(g *gin.Context) { baseAPI.CreateBaseRole(g) }
