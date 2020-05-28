package common

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// SharedK8sClient internal global use
var SharedK8sClient *K8sClientSet

// K8sClientSet interface package
type K8sClientSet struct {
	CacheInformer *dyn.CacheInformerFactory // nuwa project resource use dyn client
	ClientSetV1   *kubernetes.Clientset     // kubernetes native resource clientSet
}

// NewK8sClientSet kubernetes client required for external initialization workload
func NewK8sClientSet(cacheInformer *dyn.CacheInformerFactory, cliv1 *kubernetes.Clientset, rest *rest.Config) error {
	if SharedK8sClient != nil {
		return nil
	}
	SharedK8sClient = &K8sClientSet{
		CacheInformer: cacheInformer,
		ClientSetV1:   cliv1,
	}
	return nil
}
