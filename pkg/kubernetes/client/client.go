package client

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"

	dyn "k8s.io/client-go/dynamic"
	informers "k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdlatest "k8s.io/client-go/tools/clientcmd/api/latest"
	clientcmdapiv1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

const (
	// High enough QPS to fit all expected use cases.
	qps = 1e6
	// High enough Burst to fit all expected use cases.
	burst = 1e6
	// full resyc cache resource time
	period = 30 * time.Second
)

var SharedCacheInformerFactory *CacheInformerFactory

func buildDynamicClient(master string, config clientcmdapiv1.Config) (dyn.Interface, error) {
	cfg, err := clientcmdlatest.Scheme.ConvertToVersion(&config, clientcmdapi.SchemeGroupVersion)
	clientCfg, err := clientcmd.NewDefaultClientConfig(*(cfg.(*clientcmdapi.Config)),
		&clientcmd.ConfigOverrides{
			ClusterDefaults: clientcmdapi.Cluster{Server: master},
		}).ClientConfig()
	if err != nil {
		return nil, err
	}

	clientCfg.QPS = qps
	clientCfg.Burst = burst
	dynClient, err := dyn.NewForConfig(clientCfg)
	if err != nil {
		return nil, err
	}

	return dynClient, nil
}

type CacheInformerFactory struct {
	client   dyn.Interface
	informer informers.DynamicSharedInformerFactory
	stopChan chan struct{}
}

func NewCacheInformerFactory(master string, config clientcmdapiv1.Config) (*CacheInformerFactory, error) {
	if SharedCacheInformerFactory != nil {
		return SharedCacheInformerFactory, nil
	}
	client, err := buildDynamicClient(master, config)
	if err != nil {
		return nil, err
	}
	stop := make(chan struct{})
	sharedInformerFactory := informers.NewDynamicSharedInformerFactory(client, period)

	for _, v := range GroupVersionResources {
		go sharedInformerFactory.ForResource(v).Informer().Run(stop)
	}

	sharedInformerFactory.Start(stop)

	SharedCacheInformerFactory = &CacheInformerFactory{
		client, sharedInformerFactory, stop,
	}

	return SharedCacheInformerFactory, nil
}

func (c *CacheInformerFactory) List(res ResourceName, ns string, opt metav1.ListOptions) ([]byte, error) {
	unstructd, err := c.client.Resource(gvr(res)).Namespace(ns).List(opt)

	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(unstructd)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (c *CacheInformerFactory) Get(res ResourceName, ns, name string, opt metav1.GetOptions) ([]byte, error) {
	unstructd, err := c.client.Resource(gvr(res)).Namespace(ns).Get(name, opt)

	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(unstructd)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (c *CacheInformerFactory) Client() dyn.Interface {
	if c.client == nil {
		panic("client not initialized")
	}
	return c.client
}
