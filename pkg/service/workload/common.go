package workload

import (
	"sort"

	fv1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Resource query conditions
const (
	// ResourceNamespaceKey query resource version predicate
	ResourcePredicateNamespaceKey = "namespace" // "namespace":"default"
)

// WorkloadsSlice query resource results
type WorkloadsSlice []*fv1.Workloads

func (w WorkloadsSlice) Len() int      { return len(w) }
func (w WorkloadsSlice) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
func (w WorkloadsSlice) Less(i, j int) bool {
	if w[i].ObjectMeta.CreationTimestamp.Before(&w[j].ObjectMeta.CreationTimestamp) {
		return true
	}
	return false
}

// ResourceQuery query resource interface
type ResourceQuery interface {
	List(resource schema.GroupVersionResource, namespace, flag string, pos, size int64) (list *unstructured.UnstructuredList, err error)
	Get(resource schema.GroupVersionResource, namespace, name string) (item *unstructured.Unstructured, err error)
	Watch(resource schema.GroupVersionResource, namespace, name string) (itemChan chan *unstructured.Unstructured, closed chan struct{}, err error)
}

// ResourceApply update resource interface
type ResourceApply interface {
	Apply(resource schema.GroupVersionResource, obj *unstructured.Unstructured) error
	Delete(resource schema.GroupVersionResource, namespace, name string) error
}

// History history resource interface
type History interface {
	HistoryGet(namespace, name string) (*fv1.Workloads, error)
	HistorySave(namespace string, workloads *fv1.Workloads) error
	HistoryList(selector labels.Selector, limit int) (WorkloadsSlice, error)
}

// WorkloadsResourceHandler all needed interface defined
type WorkloadsResourceHandler interface {
	History
	ResourceQuery
	ResourceApply
}

// check the default implemented
var _ WorkloadsResourceHandler = &defaultImplWorkloadsResourceHandler{}

type defaultImplWorkloadsResourceHandler struct{}

func (d defaultImplWorkloadsResourceHandler) HistoryGet(namespace, name string) (*fv1.Workloads, error) {
	return sharedClientsCacheSet.fuxiClientSet.client.FuxiV1().Workloadses(namespace).Get(name, metav1.GetOptions{})
}

func (d defaultImplWorkloadsResourceHandler) HistorySave(namespace string, workloads *fv1.Workloads) error {
	_, err := sharedClientsCacheSet.fuxiClientSet.client.FuxiV1().Workloadses(namespace).Create(workloads)
	if err != nil {
		return err
	}
	return nil
}

// TODO 参数selector 改成string
func (d defaultImplWorkloadsResourceHandler) HistoryList(selector labels.Selector, limit int) (WorkloadsSlice, error) {
	list, err := sharedClientsCacheSet.fuxiClientSet.informer.Fuxi().V1().Workloadses().Lister().List(selector)
	if err != nil {
		return nil, err
	}
	workloads := WorkloadsSlice(list)
	sort.Sort(workloads)
	if limit > 0 && len(workloads) <= limit {
		workloads = workloads[0 : limit-1]
	}
	return workloads, nil
}

func (d defaultImplWorkloadsResourceHandler) List(resource schema.GroupVersionResource, namespace, flag string, pos, size int64) (list *unstructured.UnstructuredList, err error) {
	labelFromSet := labels.SelectorFromSet(
		labels.Set(map[string]string{
			ResourcePredicateNamespaceKey: namespace,
		}),
	)
	opt := metav1.ListOptions{
		Continue:      flag,
		LabelSelector: labelFromSet.String(),
	}
	if size > 0 {
		opt.Limit = size + pos
	}
	items, err := sharedClientsCacheSet.dynClient.Client().Resource(resource).List(opt)
	if err != nil {
		return nil, err
	}
	items.Items = items.Items[pos : pos+size]

	return items, nil
}

func (d defaultImplWorkloadsResourceHandler) Get(resource schema.GroupVersionResource, namespace, name string) (item *unstructured.Unstructured, err error) {
	opt := metav1.GetOptions{}
	item, err = sharedClientsCacheSet.dynClient.Client().Resource(resource).Get(name, opt)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (d defaultImplWorkloadsResourceHandler) Watch(resource schema.GroupVersionResource, namespace, name string) (itemChan chan *unstructured.Unstructured, closed chan struct{}, err error) {
	itemChan = make(chan *unstructured.Unstructured)
	closed = make(chan struct{})
	recv, err := sharedClientsCacheSet.dynClient.Client().Resource(resource).Watch(metav1.ListOptions{})
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

func (d defaultImplWorkloadsResourceHandler) Apply(resource schema.GroupVersionResource, obj *unstructured.Unstructured) error {
	if _, err := sharedClientsCacheSet.dynClient.Client().Resource(resource).Update(obj, metav1.UpdateOptions{}); err != nil {
		return err
	}
	return nil
}

func (d defaultImplWorkloadsResourceHandler) Delete(resource schema.GroupVersionResource, namespace, name string) error {
	if err := sharedClientsCacheSet.dynClient.Client().Resource(resource).Delete(name, &metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}
