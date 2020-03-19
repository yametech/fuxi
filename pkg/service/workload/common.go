package workload

import (
	fv1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	versioned "github.com/yametech/fuxi/pkg/client/clientset/versioned"
	"github.com/yametech/fuxi/pkg/client/informers/externalversions"
	k8sclient "github.com/yametech/fuxi/pkg/k8s/client"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
	"time"
)

const (
	LableForResourceTypeHistory = "history" // "history":"${RESOURCE_NAME}"
	LableIsLatest               = "latest"  // "latest":true
)

type WorkloadsSlice []*fv1.Workloads

func (w WorkloadsSlice) Len() int      { return len(w) }
func (w WorkloadsSlice) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
func (w WorkloadsSlice) Less(i, j int) bool {
	if w[i].ObjectMeta.CreationTimestamp.Before(&w[j].ObjectMeta.CreationTimestamp) {
		return true
	}
	return false
}

type ResourceGenerator interface {
	List(pos int64, size int64, flag string) (list *unstructured.UnstructuredList, err error)
	Get(name string, resourceVersion string) (item *unstructured.Unstructured, err error)
	Watch(name string) (itemChan chan *unstructured.Unstructured, closed chan struct{}, err error)
	Update(obj *unstructured.Unstructured) error
	Delete(name string) error
	Attach(name string) (itemChan chan *unstructured.Unstructured, close chan struct{}, err error)
}

type StoreVersions interface {
	QueryList(res, namespace, name string, limit int) (WorkloadsSlice, error)
}

type WorkloadsResourceHandler interface {
	ResourceGenerator
	StoreVersions
}

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

//
//
//func Visit(obj interface{}, res ResourceGenerator) error {
//	switch obj.(type) {
//	case *Deployment:
//	case *Stone:
//		_ = obj.(*Stone)
//	default:
//	}
//	return nil
//}

func Visit(obj interface{}, res ResourceGenerator) error {
	switch obj.(type) {
	case *Deployment:
	case *Stone:
		_ = obj.(*Stone)
	default:
	}
	return nil
}
