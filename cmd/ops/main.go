package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/yametech/fuxi/pkg/api/ops/handler"
	"github.com/yametech/fuxi/pkg/logging"
	pri "github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/common"
	"github.com/yametech/fuxi/pkg/service/ops"
	"k8s.io/sample-controller/pkg/signals"
	"time"
)

const (
	name = "go.micro.api.ops"
	ver  = "v1alpha"
)

func main() {
	service, apiInstallConfigure, err := pri.InitApi(50, name, ver, "")
	if err != nil {
		logging.Log.Error(err)
	}

	common.SharedK8sClient = &apiInstallConfigure.DefaultInstallConfigure

	tektonClient, err := tektonclient.NewForConfig(apiInstallConfigure.RestConfig)
	if err != nil {
		logging.Log.Error("Initail tekton clientv2 error", err)
	}
	ops.TektonClient = tektonClient

	router := gin.Default()
	//router.Use(gin2micro.TracerWrapper)

	group := router.Group("/ops")

	resyncDur := time.Second * 30
	op := ops.NewOps(resyncDur)
	stopCh := signals.SetupSignalHandler()
	op.Start(stopCh)

	ops := handler.OpsController{
		//Service: op,
	}

	// ops
	{
		//log
		group.GET("/v1/log/:namespace/:pipelinename", ops.GetRealLog)
		group.GET("/v1/taskrun/log/:namespace/:taskrunname", ops.GetTaskRunLog)
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
