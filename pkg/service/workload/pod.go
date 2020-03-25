package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Pod struct{ Namespace string }

func (p *Pod) List(pos int64, size int64, flag string) (list *unstructured.UnstructuredList, err error) {
	return nil, nil
}
func (p *Pod) Get(name string, resourceVersion string) (item *unstructured.Unstructured, err error) {
	return nil, nil
}

func (p *Pod) Watch(name string) (itemChan chan *unstructured.Unstructured, closed chan struct{}, err error) {
	itemChan = make(chan *unstructured.Unstructured)
	closed = make(chan struct{})

	watch, err := dyn.SharedCacheInformerFactory.Client().
		Resource(dyn.ResourcePod).
		Watch(metav1.ListOptions{Watch: true})

	if err != nil {
		return nil, nil, err
	}
	go func() {
		for {
			select {
			case <-closed:
				close(closed)
				return
			case event, ok := <-watch.ResultChan():
				if !ok {
					return
				}
				item := event.Object.(*unstructured.Unstructured)
				itemChan <- item
			}
		}
	}()
	return itemChan, closed, nil
}

func (p *Pod) Update(obj *unstructured.Unstructured) error { return nil }

func (p *Pod) Delete(name string) error { return nil }
