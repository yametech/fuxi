package main

import (
	//"log"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/yametech/fuxi/pkg/api/workload/handler"
	"github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/tracer/opentracing/gin2micro"

	// swagger doc
	file "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"
	_ "github.com/yametech/fuxi/cmd/workload/docs"
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
	name = "go.micro.api.workload"
	ver  = "latest"
)

func main() {
	service, _, err := preinstall.InitApi(50, name, ver, "")
	if err != nil {
		panic(err)
	}

	hystrix.DefaultTimeout = 5000
	wrapper := hystrixplugin.NewClientWrapper()
	_ = wrapper

	router := gin.Default()
	router.Use(gin2micro.TracerWrapper)
	router.Use()

	handler.CreateSharedSessionManager()
	workloadsAPI := &handler.WorkloadsAPI{}

	router.GET("/workload/attach", gin.WrapH(handler.CreateAttachHandler("/workload/attach")))
	router.GET("/workload/pod", workloadsAPI.PodAttach)

	/// Then, if you set envioment variable DEV_OPEN_SWAGGER to anything, /swagger/*any will respond 404, just like when route unspecified.
	/// Release production environment can be turned on
	router.GET("/workload/swagger/*any", swag.DisablingWrapHandler(file.Handler, "DEV_OPEN_SWAGGER"))

	group := router.Group("/workload")
	_ = group
	//group.GET("/list/:ns/deployment", ListDeployments)
	//group.GET("/resource", GetResource)

	service.Handle("/", router)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
