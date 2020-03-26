package main

import (
	//"log"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/yametech/fuxi/pkg/api/workload/handler"
	"github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/tracer/opentracing/gin2micro"
	// swagger doc
	file "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"
	_ "github.com/yametech/fuxi/cmd/workload/docs"
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
	name = "go.micro.api.workload"
	ver  = "latest"
)

func main() {
	service, _, err := preinstall.InitApi(50, name, ver, "")
	if err != nil {
		panic(err)
	}

	hystrix.DefaultTimeout = 5000
	wrapper := hystrixplugin.NewClientWrapper()
	_ = wrapper

	router := gin.Default()
	router.Use(gin2micro.TracerWrapper)
	router.Use()

	handler.CreateSharedSessionManager()
	workloadApi := &handler.WorkloadsApi{}
	router.GET("/workload/attach", gin.WrapH(handler.CreateAttachHandler("/workload/attach")))
	router.GET("/workload/attach/pod", workloadApi.PodAttach)

	/// Then, if you set envioment variable DEV_OPEN_SWAGGER to anything, /swagger/*any will respond 404, just like when route unspecified.
	/// Release production environment can be turned on
	router.GET("/workload/swagger/*any", swag.DisablingWrapHandler(file.Handler, "DEV_OPEN_SWAGGER"))

	group := router.Group("/workload")
	_ = group
	//group.GET("/list/:ns/deployment", ListDeployments)
	//group.GET("/resource", GetResource)

	service.Handle("/", router)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//
//type event struct {
//	pubsub.Publisher
//	clients ws.Clients
//}
//
//func newEvent() event {
//	return event{clients: ws.NewClients()}
//}
//
//func (e *event) Collection() {
//	for {
//		// call kubernetes client watch api receive data
//		e.Publish(`{"msg"":123}`)
//		time.Sleep(1 * time.Second)
//	}
//}
//
//// /workload/v1/:ns/event
//func (e *event) Event(g *gin.Context) {
//	deadline := time.Now().Add(time.Second * 2)
//	// Upgrade initial GET request to a websocket
//	conn, err := e.clients.Upgrade(g, 1024, 4096, deadline)
//	if err != nil {
//		//
//	}
//	closeChan := make(chan struct{})
//	r, _ := e.SubChannel(nil)
//	go func() {
//		for {
//			select {
//			case <-closeChan:
//				return
//			case msg, ok := <-r:
//				if !ok {
//					ws.CloseTheWebsocket(conn)
//				}
//				if err := conn.WriteJSON(msg); err != nil {
//					ws.CloseTheWebsocket(conn)
//				}
//			}
//		}
//	}()
//	// recv publish
//	for {
//		var data map[string]string
//		if err := conn.ReadJSON(data); err != nil {
//			break
//		}
//	}
//	closeChan <- struct{}{}
//}
