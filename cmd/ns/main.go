//package main

//
//import (
//	"github.com/micro/go-micro/util/log"
//
//	"github.com/micro/go-micro"
//	"github.com/yametech/fuxi/cmd/ns/client"
//)
//
//func main() {
//	// New Service
//	service := micro.NewService(
//		micro.Name("go.micro.api.ns"),
//		micro.Version("latest"),
//	)
//
//	// Initialise service
//	service.Init(
//		// create wrap for the Ns srv client
//		micro.WrapHandler(client.NsWrapper(service)),
//	)
//	//
//	//// Register Handler
//	//ns.RegisterNsHandler(service.Server(), new(handler.Ns))
//
//	// Run service
//	if err := service.Run(); err != nil {
//		log.Fatal(err)
//	}
//}

package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/yametech/fuxi/pkg/api/namespace/handler"
	pri "github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/tracer/opentracing/gin2micro"
)

const (
	name = "go.micro.api.ns"
	ver  = "v1alpha"
)

func main() {
	service, _, err := pri.InitApi(50, name, ver, "")
	if err != nil {
		log.Error(err)
	}

	hystrix.DefaultTimeout = 5000
	wrapper := hystrixplugin.NewClientWrapper()
	_ = wrapper

	router := gin.Default()
	router.Use(gin2micro.TracerWrapper)

	group := router.Group("/ns")
	hystrix.DefaultTimeout = 5000
	sClient := hystrixplugin.NewClientWrapper()(service.Options().Service.Client())
	sClient.Init(
		client.Retries(3),
	)

	ns := handler.New(sClient)
	// ns
	{

		group.POST("/v1/namespace", ns.CreateNamespace)
		group.DELETE("/v1/namespace", ns.DeleteNamespace)
		group.GET("/v1/namespacelist", ns.NamespaceList)
		group.PUT("/v1/namespaceedit", ns.EditNamespace)
	}

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
