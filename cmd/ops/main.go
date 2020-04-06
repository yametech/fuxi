package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/yametech/fuxi/pkg/api/ops/handler"
	kubeclient "github.com/yametech/fuxi/pkg/k8s/client"
	"github.com/yametech/fuxi/pkg/logging"
	pri "github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/ops"
	"github.com/yametech/fuxi/pkg/tekton"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/tracer/opentracing/gin2micro"
)

const (
	name = "go.micro.api.ops"
	ver  = "v1alpha"
)

func main() {
	service, _, err := pri.InitApi(50, name, ver, "")
	if err != nil {
		logging.Log.Error(err)
	}

	tektonClient, err := tektonclient.NewForConfig(kubeclient.RestConf)
	if err != nil {
		logging.Log.Error("Initail tekton client error", err)
	}
	tekton.TektonClient = tektonClient

	hystrix.DefaultTimeout = 5000
	wrapper := hystrixplugin.NewClientWrapper()
	_ = wrapper

	router := gin.Default()
	router.Use(gin2micro.TracerWrapper)

	group := router.Group("/ops")
	hystrix.DefaultTimeout = 5000
	sClient := hystrixplugin.NewClientWrapper()(service.Options().Service.Client())
	sClient.Init(
		client.Retries(3),
	)

	ops := handler.OpsController{
		Service: ops.NewOps(),
	}
	// ops
	{
		group.GET("/v1/log", ops.GetRealLog)
		//group.GET("/v1/repos", ops.ListRepos)
		//group.GET("/v1/branchs", ops.ListBranchs)
	}

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
