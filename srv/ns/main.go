package main

import (
	"time"

	ovnclient "github.com/alauda/kube-ovn/pkg/client/clientset/versioned"
	"github.com/golang/glog"
	"github.com/micro/go-micro/util/log"
	"github.com/yametech/fuxi/pkg/db"
	"github.com/yametech/fuxi/pkg/kubernetes/clientv1"
	"github.com/yametech/fuxi/pkg/ovn"
	pri "github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/srv/ns/handler"
	"github.com/yametech/fuxi/srv/ns/informer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	ns "github.com/yametech/fuxi/proto/ns"
)

var (
	name    = "go.micro.srv.ns"
	version = "latest"
)

func main() {
	// New Service
	service := pri.InitService(name, version)
	// Initialise service
	service.Init()

	var err error
	ovn.KubeOvnClient, err = ovnclient.NewForConfig(clientv1.RestConf)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrateNamespace()

	Run(clientv1.K8sClient)
	// Register Handler
	ns.RegisterNsHandler(service.Server(), new(handler.Ns))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// Run is run
func Run(clientset kubernetes.Interface) {
	factory := informers.NewSharedInformerFactory(clientset, time.Hour*24)

	controller := informer.NewNamespaceLoggingController(factory)
	stop := make(chan struct{})
	err := controller.Run(stop)
	if err != nil {
		glog.Fatal(err)
	}
}
