package workload

import (
	"github.com/yametech/fuxi/pkg/client/clientset/versioned"
	"github.com/yametech/fuxi/pkg/client/informers/externalversions"
	k8sclient "github.com/yametech/fuxi/pkg/k8s/client"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

var (
	workloadsGvr = schema.GroupVersionResource{Group: "fuxi.nip.io", Version: "v1", Resource: "workloads"}
)

// sharedK8sClient internal global use
var sharedK8sClient *k8sClientSet

type clientSet struct {
	stopChan chan struct{}
	client   *versioned.Clientset
	informer externalversions.SharedInformerFactory
}

func newClientSet(rest *rest.Config) (*clientSet, error) {
	stop := make(chan struct{})
	client, err := versioned.NewForConfig(rest)
	if err != nil {
		return nil, err
	}
	//informer := externalversions.NewSharedInformerFactory(client, time.Duration(time.Second*30))
	//genericInformer, err := informer.ForResource(workloadsGvr)
	//if err != nil {
	//	return nil, err
	//}
	//go genericInformer.Informer().Run(stop)
	//
	//informer.Start(stop)

	return &clientSet{
		stopChan: stop,
		client:   client,
		//informer: informer,
	}, nil
}

// k8sClientSet interface package
type k8sClientSet struct {
	cacheInformer   *dyn.CacheInformerFactory  // nuwa project resource use dyn client
	resourceHandler *k8sclient.ResourceHandler // kubernetes native resource clients
	clientSet       *clientSet                 // fuxi resoruce client
}

// NewK8sClientSet kubernetes client required for external initialization workload
func NewK8sClientSet(cacheInformer *dyn.CacheInformerFactory, res *k8sclient.ResourceHandler, rest *rest.Config) error {
	if sharedK8sClient != nil {
		return nil
	}
	clientSet, err := newClientSet(rest)
	if err != nil {
		return err
	}
	sharedK8sClient = &k8sClientSet{
		cacheInformer:   cacheInformer,
		resourceHandler: res,
		clientSet:       clientSet,
	}
	return nil
}
