package main

import (
	"github.com/micro/micro/cmd"
	"github.com/yametech/fuxi/pkg/api/gateway/handler"
	"github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/common"
)

const name = "API gateway"

func main() {
	loginHandler := &handler.LoginHandle{Authorization: handler.NewAuthorization()}
	gatewayInstallConfigure, err := preinstall.InitGatewayInstallConfigure(name, loginHandler.Check, loginHandler)
	if err != nil {
		panic(err)
	}
	common.SharedK8sClient = &gatewayInstallConfigure.DefaultInstallConfigure

	cmd.Init()
}
