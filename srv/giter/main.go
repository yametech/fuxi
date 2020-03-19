package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/yametech/fuxi/srv/giter/handler"
	"github.com/yametech/fuxi/srv/giter/subscriber"

	giter "github.com/yametech/fuxi/srv/giter/proto/giter"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.giter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	giter.RegisterGiterHandler(service.Server(), new(handler.Giter))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.giter", service.Server(), new(subscriber.Giter))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.giter", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
