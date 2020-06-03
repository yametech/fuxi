package main

import (
	"github.com/micro/micro/cmd"
	"github.com/yametech/fuxi/pkg/api/gateway/handler"
	"github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/common"
)

const name = "API gateway"

func main() {
	loginHandler := &handler.LoginHandle{}
	gatewayInstallConfigure, err := preinstall.InitGatewayInstallConfigure(name, loginHandler)
	if err != nil {
		panic(err)
	}
	common.SharedK8sClient = &gatewayInstallConfigure.DefaultInstallConfigure
	authorStorage, err := handler.NewAuthorizationStorage()
	if err != nil {
		panic(err)
	}
	loginHandler.IAuthorStorage = authorStorage

	cmd.Init()
}
