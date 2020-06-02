package clientv2

import (
	clientv2 "k8s.io/client-go/dynamic"
	informers "k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdlatest "k8s.io/client-go/tools/clientcmd/api/latest"
	clientcmdapiv1 "k8s.io/client-go/tools/clientcmd/api/v1"
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"time"
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

func buildDynamicClient(master string, config clientcmdapiv1.Config) (clientv2.Interface, error) {
	cfg, err := clientcmdlatest.Scheme.ConvertToVersion(&config, clientcmdapi.SchemeGroupVersion)
	clientCfg, err := clientcmd.NewDefaultClientConfig(*(cfg.(*clientcmdapi.Config)),
		&clientcmd.ConfigOverrides{ClusterDefaults: clientcmdapi.Cluster{Server: master}}).ClientConfig()
	if err != nil {
		return nil, err
	}

	clientCfg.QPS = qps
	clientCfg.Burst = burst
	dynClient, err := clientv2.NewForConfig(clientCfg)
	if err != nil {
		return nil, err
	}

	return dynClient, nil
}

type CacheInformerFactory struct {
	Interface clientv2.Interface
	Informer  informers.DynamicSharedInformerFactory
	stopChan  chan struct{}
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

	for _, v := range types.GroupVersionResources {
		go sharedInformerFactory.ForResource(v).Informer().Run(stop)
	}

	sharedInformerFactory.Start(stop)

	SharedCacheInformerFactory = &CacheInformerFactory{
		client, sharedInformerFactory, stop,
	}

	return SharedCacheInformerFactory, nil
}
