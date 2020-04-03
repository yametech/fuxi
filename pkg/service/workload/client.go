package workload

import (
	fv1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/client/clientset/versioned"
	"github.com/yametech/fuxi/pkg/client/informers/externalversions"
	k8sclient "github.com/yametech/fuxi/pkg/k8s/client"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/client-go/rest"
	"time"
)

var SharedClientsCacheSet *ClientsCacheSet

type FuxiClientSet struct {
	stopChan              chan struct{}
	SharedInformerFactory externalversions.SharedInformerFactory
}

func NewFuxiClientSet(rest *rest.Config) (*FuxiClientSet, error) {
	stop := make(chan struct{})
	client, err := versioned.NewForConfig(rest)
	if err != nil {
		return nil, err
	}
	sharedInformerFactory := externalversions.NewSharedInformerFactory(client, time.Duration(time.Second*30))
	genericInformer, err := sharedInformerFactory.ForResource(fv1.SchemeGroupVersion.WithResource("workloads"))
	if err != nil {
		return nil, err
	}
	go genericInformer.Informer().Run(stop)

	sharedInformerFactory.Start(stop)

	return &FuxiClientSet{
		stopChan:              stop,
		SharedInformerFactory: sharedInformerFactory,
	}, nil
}

type ClientsCacheSet struct {
	dynClient     *dyn.CacheInformerFactory  // nuwa resource client
	defaultClient *k8sclient.ResourceHandler // kubernetes native resource clients
	fxClientSet   *FuxiClientSet             // fuxi resoruce client
}

func NewClientsCacheSet(dynClient *dyn.CacheInformerFactory, defaultClient *k8sclient.ResourceHandler, rest *rest.Config) (*ClientsCacheSet, error) {
	fuxiClientSet, err := NewFuxiClientSet(rest)
	if err != nil {
		return nil, err
	}
	return &ClientsCacheSet{
		dynClient:     dynClient,
		defaultClient: defaultClient,
		fxClientSet:   fuxiClientSet,
	}, nil
}
