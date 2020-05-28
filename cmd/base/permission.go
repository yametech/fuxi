package main

import (
	"github.com/gin-gonic/gin"
)

// BasePermission doc
// @Summary base permission list
// @Description base service for list base permission
// @Tags BasePermission
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/basepermissions [get]
func BasePermissionList(g *gin.Context) { baseAPI.ListBasePermission(g) }

// BasePermission doc
// @Summary base permission get
// @Description base service for get a base permission detail
// @Tags BasePermission
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/basepermissions/{name} [get]
func BasePermissionGet(g *gin.Context) { baseAPI.GetBasePermission(g) }

// BasePermission doc
// @Summary base permission list
// @Description base service for base permission
// @Tags BasePermission
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/basepermissions/{name} [post]
func BasePermissionCreate(g *gin.Context) { baseAPI.CreateBasePermission(g) }
