package main

import (
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/yametech/fuxi/pkg/preinstall"
	"log"
	"net"
	"net/http"
)

const name = "API gateway"

func main() {
	gateWayInstall, err := preinstall.InitGateWayInstall(cors.NewPlugin())
	if err != nil {
		panic(err)
	}

	_ = gateWayInstall

	go func() {
		err := http.ListenAndServe(net.JoinHostPort("", "8080"), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	cmd.Init()
}
