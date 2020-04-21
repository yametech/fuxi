package workload

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sort"

	fv1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
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
	List(resource schema.GroupVersionResource, namespace, flag string, pos, size int64, selector interface{}) (*unstructured.UnstructuredList, error)
	Get(resource schema.GroupVersionResource, namespace, name string) (runtime.Object, error)
	Watch(resource schema.GroupVersionResource, namespace string, resourceVersion string, timeoutSeconds int64, selector labels.Selector, closed chan struct{}) (<-chan watch.Event, error)
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
	HistoryList(namespace string, selector labels.Selector, limit int) (WorkloadsSlice, error)
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

func (d *defaultImplWorkloadsResourceHandler) HistoryGet(namespace, name string) (*fv1.Workloads, error) {
	return sharedK8sClient.
		clientSet.
		client.
		FuxiV1().
		Workloadses(namespace).
		Get(name, metav1.GetOptions{})
}

func (d *defaultImplWorkloadsResourceHandler) HistorySave(namespace string, workloads *fv1.Workloads) error {
	_, err := sharedK8sClient.
		clientSet.
		client.
		FuxiV1().
		Workloadses(namespace).
		Create(workloads)
	if err != nil {
		return err
	}
	return nil
}

func (d *defaultImplWorkloadsResourceHandler) HistoryList(
	namespace string,
	selector labels.Selector,
	limit int,
) (WorkloadsSlice, error) {
	list, err := sharedK8sClient.
		clientSet.
		informer.
		Fuxi().
		V1().
		Workloadses().
		Lister().
		Workloadses(namespace).
		List(selector)
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

func (d *defaultImplWorkloadsResourceHandler) List(
	resource schema.GroupVersionResource,
	namespace,
	flag string,
	pos,
	size int64,
	selector interface{},
) (*unstructured.UnstructuredList, error) {
	var err error
	var items *unstructured.UnstructuredList
	opts := metav1.ListOptions{}

	if selector == nil || selector == "" {
		selector = labels.Everything()
	}
	switch selector.(type) {
	case labels.Selector:
		opts.LabelSelector = selector.(labels.Selector).String()
	case string:
		if selector != "" {
			opts.LabelSelector = selector.(string)
		}
	}

	if flag != "" {
		opts.Continue = flag
	}
	if size > 0 {
		opts.Limit = size + pos
	}
	items, err = sharedK8sClient.
		cacheInformer.
		Client.
		Resource(resource).
		Namespace(namespace).
		List(opts)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (d *defaultImplWorkloadsResourceHandler) Get(
	resource schema.GroupVersionResource,
	namespace,
	name string,
) (runtime.Object, error) {
	object, err := sharedK8sClient.
		cacheInformer.
		Informer.
		ForResource(resource).
		Lister().
		ByNamespace(namespace).Get(name)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (d *defaultImplWorkloadsResourceHandler) Watch(
	resource schema.GroupVersionResource,
	namespace string,
	resourceVersion string,
	timeoutSeconds int64,
	selector labels.Selector,
	closed chan struct{},
) (<-chan watch.Event, error) {
	closed = make(chan struct{})
	opts := metav1.ListOptions{}
	if selector != nil {
		opts.LabelSelector = selector.String()
	}
	if timeoutSeconds > 0 {
		opts.TimeoutSeconds = &timeoutSeconds
	}
	if resourceVersion != "" {
		opts.ResourceVersion = resourceVersion
	}
	recv, err := sharedK8sClient.
		cacheInformer.
		Client.
		Resource(resource).
		Namespace(namespace).
		Watch(opts)
	if err != nil {
		return nil, err
	}
	return recv.ResultChan(), nil
}

func (d *defaultImplWorkloadsResourceHandler) Apply(
	resource schema.GroupVersionResource,
	obj *unstructured.Unstructured,
) error {
	_, err := sharedK8sClient.
		cacheInformer.
		Client.
		Resource(resource).
		Update(obj, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (d *defaultImplWorkloadsResourceHandler) Delete(
	resource schema.GroupVersionResource,
	namespace, name string,
) error {
	err := sharedK8sClient.
		cacheInformer.
		Client.
		Resource(resource).
		Namespace(namespace).
		Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
