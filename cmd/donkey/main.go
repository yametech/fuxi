package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/common"
	"github.com/yametech/fuxi/pkg/service/donkey"
)

const (
	name = "go.micro.srv.donkey"
	ver  = "v1"
)

func initNeed() micro.Service {
	service, apiInstallConfigure := preinstall.InitService(name, ver)
	common.SharedK8sClient = &apiInstallConfigure.DefaultInstallConfigure

	return service
}

func main() {
	service := initNeed()
	var assistant donkey.IAssistant = donkey.NewDepartmentAssistant()

	go func() {
		if err := assistant.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := service.Run(); err != nil {
		assistant.Stop()
		log.Fatal(err)
	}
}
