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
	"k8s.io/sample-controller/pkg/signals"
	"time"
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

	resyncDur := time.Second * 30
	op := ops.NewOps(resyncDur)
	stopCh := signals.SetupSignalHandler()
	op.Start(stopCh)

	ops := handler.OpsController{
		Service: op,
	}

	// ops
	{
		//log
		group.GET("/v1/log/:namespace/:pipelinename", ops.GetRealLog)
		group.GET("/v1/taskrun/log/:namespace/:taskrunname",ops.GetTaskRunLog)
		//pipeline
		group.POST("/v1/pipeline", ops.CreateOrUpdatePipeline)
		group.GET("/v1/pipelines/:namespace", ops.PipelineList)
		group.GET("/v1/pipeline/:namespace/:name", ops.GetPipeline)
		group.DELETE("/v1/pipeline/:namespace", ops.PipelineDelete)
		//task
		group.POST("/v1/task/", ops.CreateOrUpdateTask)
		group.GET("/v1/tasks/:namespace", ops.TaskList)
		group.GET("/v1/task/:namespace/:name", ops.GetTask)
		group.DELETE("/v1/task/:namespace/:name", ops.DeleteTask)
		//resource
		group.POST("/v1/resource/", ops.CreateOrUpdatePipelineResource)
		group.GET("/v1/resources/:namespace", ops.PipelineResourceList)
		group.GET("/v1/resource/:namespace/:name", ops.GetPipelineResource)
		group.DELETE("/v1/resource/:namespace/:name", ops.DeleteTask)
		//pipelinerun
		group.POST("/v1/pipelinerun", ops.CreateOrUpdatePipelineRun)
		group.GET("/v1/pipelineruns/latest/:namespace", ops.GetLatestPipelineRunList)
		group.GET("/v1/pipelineruns/history/:namespace/:name", ops.GetPipelineRunHistoryList)
		group.DELETE("/v1/pipelinerun/:namespace/:name", ops.PipelineRunDelete)
		group.GET("/v1/pipelinerun/:namespace/:name", ops.GetPipelineRun)
		group.POST("/v1/pipelinerun/rerun/:namespace/:name", ops.ReRunPipeline)
		group.POST("/v1/pipelinerun/cancel/:namespace/:name", ops.CancelPipelineRun)
		//repos
		group.GET("/v1/repos/:namespace", ops.ListRepos)
		group.GET("/v1/branchs/:namespace", ops.ListBranchs)
	}

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
