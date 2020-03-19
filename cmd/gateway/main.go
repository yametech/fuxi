package main

import (
	"github.com/micro/micro/cmd"
	"github.com/yametech/fuxi/pkg/preinstall"
)

const name = "API gateway"

func main() {
	preinstall.InitGateWay()

	// stdhttp.SetSamplingFrequency(50)
	// t, io, err := tracer.NewTracer(name, "")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer func() {
	// 	if err := io.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// opentracing.SetGlobalTracer(t)

	// hystrixStreamHandler := ph.NewStreamHandler()
	// hystrixStreamHandler.Start()

	// go func() {
	// 	if err := http.ListenAndServe(
	// 		net.JoinHostPort("", "18081"), hystrixStreamHandler); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	cmd.Init()
}
