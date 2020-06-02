package main

import (
	"github.com/gin-gonic/gin"
	"k8s.io/sample-controller/pkg/signals"

	"github.com/micro/go-micro/util/log"
	"github.com/yametech/fuxi/pkg/api/namespace/handler"
	pri "github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/ns"
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

	router := gin.Default()

	group := router.Group("/ns")

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
