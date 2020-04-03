package handler

import (
	cli "github.com/yametech/fuxi/pkg/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

type WorkloadInterface interface {
	Get(namespace, name string) (*unstructured.Unstructured, error)
	List(namespace string) ([]*unstructured.Unstructured, error)
	Create(object runtime.Object) error
}

type Pod struct {
}

func (p *Pod) Get(namespace, name string) (*unstructured.Unstructured, error) {
	return cli.SharedCacheInformerFactory.
		Client().Resource(cli.ResourcePod).
		Namespace(namespace).
		Get(name, metav1.GetOptions{})
}

func (p *Pod) List(namespace string) (*unstructured.UnstructuredList, error) {
	return cli.SharedCacheInformerFactory.
		Client().Resource(cli.ResourcePod).
		Namespace(namespace).
		List(metav1.ListOptions{})
}

func (p *Pod) Watch(namespace string) (watch.Interface, error) {
	return cli.SharedCacheInformerFactory.
		Client().Resource(cli.ResourcePod).
		Namespace(namespace).Watch(metav1.ListOptions{})
}

func example() {
	//pod := &Pod{}
	//
	//w, err := pod.Watch("xxx")
	//if err != nil {
	//	panic(err)
	//}
	//
	//go func() {
	//	for {
	//		select {
	//		case event, ok := <-w.ResultChan():
	//			if !ok {
	//				return
	//			}
	//			//event.Object().(*corev1.Pod)
	//			event.Object(*corev1.Pod)
	//			//ws.send()
	//		}
	//	}
	//}()
	//

}
