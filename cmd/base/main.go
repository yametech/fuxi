package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/yametech/fuxi/pkg/api/base/handler"
	"github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/common"

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

func initNeed() (web.Service, *gin.Engine, *gin.RouterGroup, *handler.BaseAPI) {
	service, apiInstallConfigure, err := preinstall.InitApi(50, name, ver, "")
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	common.SharedK8sClient = &apiInstallConfigure.DefaultInstallConfigure
	return service, router, router.Group("/base"), handler.NewBaseAPi()
}

var service, router, group, baseAPI = initNeed()

func main() {

	// BaseDepartment
	{
		group.GET("/apis/fuxi.nip.io/v1/basedepartments", BaseDepartmentList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/basedepartments/:name", BaseDepartmentGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/basedepartments", BaseDepartmentCreate)
	}

	// BaseRole
	{
		group.GET("/apis/fuxi.nip.io/v1/baseroles", BaseRoleList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/baseroles/:name", BaseRoleGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/baseroles", BaseRoleCreate)
	}

	// BaseUser
	{
		group.GET("/apis/fuxi.nip.io/v1/baseusers", BaseUserList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/baseusers/:name", BaseUserGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/baseusers", BaseUserCreate)
	}

	// BaseRoleUser
	{
		group.GET("/apis/fuxi.nip.io/v1/baseroleusers", BaseRoleUserList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/baseroleusers/:name", BaseRoleUserGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/baseroleusers", BaseRoleUserCreate)
	}

	// Permission
	{
		group.GET("/permission_transfer/:value", BasePermissionTransferList)
	}

	// Release production environment can be turned on
	router.GET("/base/swagger/*any", swag.DisablingWrapHandler(file.Handler, "NAME_OF_ENV_VARIABLE"))

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
