package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	_ "github.com/yametech/fuxi/cmd/base/docs"
	"github.com/yametech/fuxi/pkg/api/workload/handler"
	"github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/common"
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
	name = "go.micro.api.webhook"
	ver  = "v1"
)

func initNeed() (web.Service, *gin.Engine, *gin.RouterGroup, *handler.WorkloadsAPI) {
	service, apiInstallConfigure, err := preinstall.InitApi(50, name, ver, "")
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	common.SharedK8sClient = &apiInstallConfigure.DefaultInstallConfigure
	return service, router, router.Group("/webhook"), handler.NewWorkladAPI()
}

var service, router, group, workloadsAPI = initNeed()

func main() {

	{
		group.POST("/gitea/namespaces/:namespace/tektonwebhooks/:name", TriggerGiteaWebHook)
	}

	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
