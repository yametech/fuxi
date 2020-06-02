package main

import "github.com/gin-gonic/gin"

// BaseDepartment doc
// @Summary base department list
// @Description base service for list base department
// @Tags BaseDepartment
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/basedepartments [get]
func BaseDepartmentList(g *gin.Context) { baseAPI.ListBaseDepartment(g) }

// BaseDepartment doc
// @Summary base department get
// @Description base service for get a base department detail
// @Tags BaseDepartment
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Param name query string true "name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/basedepartments/{name} [get]
func BaseDepartmentGet(g *gin.Context) { baseAPI.GetBaseDepartment(g) }

// BaseDepartment doc
// @Summary base department list
// @Description base service for base department
// @Tags BaseDepartment
// @Accept mpfd
// @Produce json
// @Param namespace query string true "namespace"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/apis/fuxi.nip.io/v1/namespaces/{namespace}/basedepartments/ [post]
func BaseDepartmentCreate(g *gin.Context) { baseAPI.CreateBaseDepartment(g) }
