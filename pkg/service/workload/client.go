package workload

import (
	"time"

	fv1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/client/clientset/versioned"
	"github.com/yametech/fuxi/pkg/client/informers/externalversions"
	k8sclient "github.com/yametech/fuxi/pkg/k8s/client"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/client-go/rest"
)

// sharedClientsCacheSet internal global use
var sharedClientsCacheSet *ClientsCacheSet

type fuxiClientSet struct {
	stopChan chan struct{}
	client   *versioned.Clientset
	informer externalversions.SharedInformerFactory
}

func newFuxiClientSet(rest *rest.Config) (*fuxiClientSet, error) {
	stop := make(chan struct{})
	client, err := versioned.NewForConfig(rest)
	if err != nil {
		return nil, err
	}
	informer := externalversions.NewSharedInformerFactory(client, time.Duration(time.Second*30))
	genericInformer, err := informer.ForResource(fv1.SchemeGroupVersion.WithResource("workloads"))
	if err != nil {
		return nil, err
	}
	go genericInformer.Informer().Run(stop)

	informer.Start(stop)

	return &fuxiClientSet{
		stopChan: stop,
		client:   client,
		informer: informer,
	}, nil
}

// ClientsCacheSet interface package
type ClientsCacheSet struct {
	dynClient     *dyn.CacheInformerFactory  // nuwa project resource use dyn client
	defaultClient *k8sclient.ResourceHandler // kubernetes native resource clients
	fuxiClientSet *fuxiClientSet             // fuxi resoruce client
}

// NewClientsCacheSet kubernetes client required for external initialization workload
func NewClientsCacheSet(dynClient *dyn.CacheInformerFactory, defaultClient *k8sclient.ResourceHandler, rest *rest.Config) (*ClientsCacheSet, error) {
	if sharedClientsCacheSet != nil {
		return sharedClientsCacheSet, nil
	}
	clientSet, err := newFuxiClientSet(rest)
	if err != nil {
		return nil, err
	}
	sharedClientsCacheSet = &ClientsCacheSet{
		dynClient:     dynClient,
		defaultClient: defaultClient,
		fuxiClientSet: clientSet,
	}

	return sharedClientsCacheSet, nil
}
