package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"k8s.io/sample-controller/pkg/signals"

	"github.com/micro/go-micro/util/log"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/yametech/fuxi/pkg/api/namespace/handler"
	pri "github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/ns"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/tracer/opentracing/gin2micro"
	"time"
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
	//hystrix.DefaultTimeout = 5000
	//sClient := hystrixplugin.NewClientWrapper()(service.Options().Service.Interface())
	//sClient.Init(
	//	clientv2.Retries(3),
	//)
	//
	//ns := handler.New(sClient)
	// ns
	resyncDury := time.Second * 30
	op := ns.NewNS(resyncDury)
	stopCh := signals.SetupSignalHandler()
	op.Start(stopCh)

	ns := handler.NSController{
		Service: op,
	}
	{

		group.POST("/v1/namespace", ns.CreateSubNet)
		group.DELETE("/v1/namespace/:name", ns.DeleteSubNet)
		group.GET("/v1/namespace/:name", ns.GetSubNet)
		group.PUT("/v1/namespace", ns.UpdateSubNet)
		group.GET("/v1/namespaces", ns.SubNetList)
	}

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
