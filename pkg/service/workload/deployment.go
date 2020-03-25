package workload

import (
	"fmt"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
)

var _ ResourceGenerator = &Deployment{""}

type Deployment struct{ Namespace string }

func NewDeploymentEntity(Namespace string) ResourceGenerator { return &Deployment{Namespace: Namespace} }

func (d *Deployment) QueryList(res, namespace, name string, limit int) (WorkloadsSlice, error) {
	selector, err := labels.Parse(fmt.Sprintf("%s=%s", LableForResourceTypeHistory, res))
	if err != nil {
		return nil, err
	}
	return SharedClientsCacheSet.fxClientSet.SharedInformerFactory.
		Fuxi().
		V1().
		Workloadses().
		Lister().
		Workloadses(namespace).
		List(selector)
}

func (d *Deployment) Update(obj *unstructured.Unstructured) error {
	if _, err := dyn.SharedCacheInformerFactory.Client().
		Resource(dyn.ResourceDeployment).
		Update(obj, metav1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (d *Deployment) Delete(name string) error {
	if err := dyn.SharedCacheInformerFactory.Client().
		Resource(dyn.ResourceDeployment).
		Delete(name, &metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (d *Deployment) Attach(name string) (itemChan chan *unstructured.Unstructured, close chan struct{}, err error) {
	panic("don't not implement me")
}

func (d *Deployment) Get(name string, resourceVersion string) (*unstructured.Unstructured, error) {
	opt := metav1.GetOptions{}
	if resourceVersion != "" {
		opt.ResourceVersion = resourceVersion
	}
	item, err := dyn.SharedCacheInformerFactory.
		Client().
		Resource(dyn.ResourceDeployment).
		Get(name, opt)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (d *Deployment) List(pos int64, size int64, flag string) (*unstructured.UnstructuredList, error) {
	opt := metav1.ListOptions{Continue: flag}
	if size > 0 {
		opt.Limit = size + pos
	}
	items, err := dyn.SharedCacheInformerFactory.
		Client().
		Resource(dyn.ResourceDeployment).
		List(opt)
	if err != nil {
		return nil, err
	}
	items.Items = items.Items[pos : pos+size]
	return items, nil
}

func (d *Deployment) Watch(name string) (itemChan chan *unstructured.Unstructured, closed chan struct{}, err error) {
	itemChan = make(chan *unstructured.Unstructured)
	closed = make(chan struct{})
	recv, err := dyn.SharedCacheInformerFactory.
		Client().
		Resource(dyn.ResourceDeployment).
		Watch(metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}
	go func() {
		for {
			select {
			case <-closed:
				return
			case event, ok := <-recv.ResultChan():
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
