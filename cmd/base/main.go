package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/yametech/fuxi/pkg/api/base/handler"
	"github.com/yametech/fuxi/pkg/db"
	pri "github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/tracer/opentracing/gin2micro"

	// swagger doc
	file "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"
	_ "github.com/yametech/fuxi/cmd/base/docs"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger base
// @contact.name laik author
// @contact.url  github.com/yametech
// @contact.email laik.lj@me.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
const (
	name = "go.micro.api.base"
	ver  = "v1"
)

var userApiService *handler.UserApiService
var departmentApiService *handler.DepartmentApiService
var permissionApiService *handler.PermissionApiService
var roleApiService *handler.RoleApiService

// User 用户 >>>>

// User info doc
// @Summary base User authentication
// @Description 用户登录验证
// @Tags User
// @Accept mpfd
// @Produce json
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} string "Success 成功"
// @Router /base/v1/user/login [post]
func uerAuthorization(c *gin.Context) { userApiService.UserAuthorization(c) }

// User info doc
// @Summary base User Info
// @Description 用户信息
// @Tags User
// @Accept mpfd
// @Produce json
// @Param x-auth-username header string false "JWT header"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Please login"
// @Router /base/v1/user [get]
func userInfo(c *gin.Context) { userApiService.UserInfo(c) }

// User info doc
// @Summary base User info registration
// @Description 用户注册
// @Tags User
// @Accept mpfd
// @Produce json
// @Param x-auth-username header string false "JWT header"
// @Param body body db.User true "User Model"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Failed"
// @Router /base/v1/user [post]
func userRegister(c *gin.Context) { userApiService.UserRegister(c) }

// User info doc
// @Summary base User info delete
// @Description 用户删除
// @Tags User
// @Accept mpfd
// @Produce json
// @Param x-auth-username header string false "JWT header"
// @Param ID path string true "ID"
// @Router /base/v1/user/{ID} [delete]
func userDelete(c *gin.Context) { userApiService.UserDelete(c) }

// User 用户 <<<<
// Department 部门 >>>>

// Department doc
// @Summary base department create (创建部门信息)
// @Description base services Department info create 创建部门
// @Tags Department
// @Accept mpfd
// @Produce json
// @Param body body db.Department true "Department Model"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/department [patch]
func departmentCreate(c *gin.Context) { departmentApiService.CreateDepartment(c) }

// Department doc
// @Summary base department edit (编辑部门信息)
// @Description Department info edit 编辑部门
// @Tags Department
// @Accept mpfd
// @Produce json
// @Param ID path string true "ID"
// @Param body body db.Department true "Department Model"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/department/{ID} [patch]
func departmentEdit(c *gin.Context) { departmentApiService.EditDepartment(c) }

// Department doc
// @Summary base department delete (删除部门信息)
// @Description Department info delete 删除部门
// @Tags Department
// @Accept mpfd
// @Produce json
// @Param ID path string true "ID"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/department/{ID} [delete]
func departmentDelete(c *gin.Context) { departmentApiService.DeleteDepartment(c) }

// Department doc
// @Summary base department list (部门列表)
// @Description Department info list 部门列表
// @Tags Department
// @Accept mpfd
// @Produce json
// @Param page query int true "page"
// @Param pageSize query int true "pageSize"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/departments [get]
func departmentList(c *gin.Context) { departmentApiService.DepartmentList(c) }

// Department 部门 <<<<
// Permission 权限 >>>>

// Permission doc
// @Summary base permission create (创建权限)
// @Description base services Permission info create 创建权限
// @Tags Permission
// @Accept mpfd
// @Produce json
// @Param body body db.Permission true "Permission Model"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/permission [post]
func permissionCreate(c *gin.Context) { permissionApiService.CreatePermission(c) }

// Permission doc
// @Summary base permission edit (编辑权限信息)
// @Description base services Permission info edit 编辑权限
// @Tags Permission
// @Accept mpfd
// @Produce json
// @Param ID path string true "ID"
// @Param body body db.Permission true "Permission Model"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/department/{ID} [patch]
func permissionEdit(c *gin.Context) { permissionApiService.EditPermission(c) }

// Permission doc
// @Summary base permission delete (删除权限信息)
// @Description Permission info delete 删除权限
// @Tags Permission
// @Accept mpfd
// @Produce json
// @Param ID path string true "ID"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/permission/{ID} [delete]
func permissionDelete(c *gin.Context) { permissionApiService.DeletePermission(c) }

// Permission doc
// @Summary base permission list (权限列表)
// @Description Permission info list (权限列表)
// @Tags Permission
// @Accept mpfd
// @Produce json
// @Param x-auth-username header string false "JWT header"
// @Param page query int true "page"
// @Param pageSize query int true "pageSize"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/permissions [get]
func permissionList(c *gin.Context) { permissionApiService.PermissionList(c) }

// Permission doc
// @Summary base permission list (权限详情)
// @Description Permission info list (权限详情)
// @Tags Permission
// @Accept mpfd
// @Produce json
// @Param x-auth-username header string false "JWT header"
// @Param ID path string true "ID"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/permission/{ID} [get]
func permissionDetail(c *gin.Context) { permissionApiService.DetailPermission(c) }

// Permission 权限 <<<<
// Role 角色 >>>>

// Role doc
// @Summary base role create (创建角色)
// @Description base services Role info create 创建角色
// @Tags Role
// @Accept mpfd
// @Produce json
// @Param name query string true "dept_name"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/role [post]
func roleCreate(c *gin.Context) { roleApiService.CreateRole(c) }

// Role doc
// @Summary base role edit (编辑角色信息)
// @Description base services Role info edit 编辑角色
// @Tags Role
// @Accept mpfd
// @Produce json
// @Param ID path string true "ID"
// @Param body body db.Role true "Permission Model"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/role/{ID} [patch]
func roleEdit(c *gin.Context) { roleApiService.EditRole(c) }

// Role doc
// @Summary base role delete (删除角色信息)
// @Description base services  Role info delete 删除角色
// @Tags Role
// @Accept mpfd
// @Produce json
// @Param x-auth-username header string false "JWT header"
// @Param ID path string true "ID"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/role/{ID} [delete]
func roleDelete(c *gin.Context) { roleApiService.DeleteRole(c) }

// Role doc
// @Summary base role list (角色列表)
// @Description base services  Role info list 角色列表
// @Tags Role
// @Accept mpfd
// @Produce json
// @Param x-auth-username header string false "JWT header"
// @Param page query int true "page"
// @Param pageSize query int true "pageSize"
// @Success 200 {string} string "{"msg": "Success"}"
// @Failure 400 {string} string "{"msg": "Failed"}"
// @Router /base/v1/roles [get]
func roleList(c *gin.Context) { roleApiService.RoleList(c) }

// Role 角色 <<<<

func main() {
	service, token2, err := pri.InitApi(50, name, ver, "")
	if err != nil {
		log.Error(err)
	}

	// automatic migration
	db.AutoMigrateUser()
	db.AutoMigrateDepartment()
	db.AutoMigratePermission()
	db.AutoMigrateRole()
	db.AutoMigrateRoleGroup()

	// setting wrapper
	hystrix.DefaultTimeout = 5000
	wrapper := hystrixplugin.NewClientWrapper()
	_ = wrapper

	router := gin.Default()
	router.Use(gin2micro.TracerWrapper)
	router.Use()

	group := router.Group("/base")

	userApiService = handler.NewUserApiService(nil, token2, wrapper)
	{
		// User
		group.POST("/v1/user/login", uerAuthorization)
		group.GET("/v1/user", userInfo)
		group.POST("/v1/user", userRegister)
		group.DELETE("v1/user/:ID", userDelete)
	}

	{
		// Department
		group.POST("/v1/department", departmentCreate)
		group.PATCH("/v1/department/:ID", departmentEdit)
		group.DELETE("/v1/department/:ID", departmentDelete)
		group.GET("/v1/departments", departmentList)
	}

	{
		//Permission
		group.POST("/v1/permission", permissionCreate)
		group.PATCH("/v1/permission/:ID", permissionEdit)
		group.DELETE("/v1/permission/:ID", permissionDelete)
		group.GET("/v1/permissions", permissionList)
		group.GET("/v1/permission/:ID", permissionDetail)
	}

	{
		// Role
		group.POST("/v1/role", roleCreate)
		group.PATCH("/v1/role/:ID", roleEdit)
		group.DELETE("/v1/role/:ID", roleDelete)
		group.GET("/v1/roles", roleList)
	}

	// Then, if you set envioment variable NAME_OF_ENV_VARIABLE to anything, /swagger/*any will respond 404, just like when route unspecified.
	// Release production environment can be turned on
	router.GET("/base/swagger/*any", swag.DisablingWrapHandler(file.Handler, "NAME_OF_ENV_VARIABLE"))

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
