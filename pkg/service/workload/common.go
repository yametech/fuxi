package workload

import (
	"encoding/json"
	fv1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/util/retry"
	"reflect"
	"sort"
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
	List(namespace, flag string, pos, size int64, selector interface{}) (*unstructured.UnstructuredList, error)
	Get(namespace, name string) (runtime.Object, error)
	RemoteGet(namespace, name string, subresources ...string) (runtime.Object, error)
	Watch(namespace string, resourceVersion string, timeoutSeconds int64, selector labels.Selector) (<-chan watch.Event, error)
}

// ResourceApply update resource interface
type ResourceApply interface {
	Apply(namespace, name string, obj *unstructured.Unstructured) (*unstructured.Unstructured, error)
	Patch(namespace, name string, patchData map[string]interface{}) (*unstructured.Unstructured, error)
	Delete(namespace, name string) error
}

// History history resource interface
type History interface {
	HistoryGet(namespace, name string) (*fv1.Workloads, error)
	HistorySave(namespace string, workloads *fv1.Workloads) error
	HistoryList(namespace string, selector labels.Selector, limit int) (WorkloadsSlice, error)
}

type WorkloadsResourceVersion interface {
	SetGroupVersionResource(schema.GroupVersionResource)
	GetGroupVersionResource() schema.GroupVersionResource
}

// WorkloadsResourceHandler all needed interface defined
type WorkloadsResourceHandler interface {
	History
	ResourceQuery
	ResourceApply
	WorkloadsResourceVersion
}

// check the default implemented
var _ WorkloadsResourceHandler = &defaultImplWorkloadsResourceHandler{}

type defaultImplWorkloadsResourceHandler struct {
	groupVersionResource schema.GroupVersionResource
}

func (d *defaultImplWorkloadsResourceHandler) GetGroupVersionResource() schema.GroupVersionResource {
	return d.groupVersionResource
}
func (d *defaultImplWorkloadsResourceHandler) SetGroupVersionResource(g schema.GroupVersionResource) {
	d.groupVersionResource = g
}

func (d *defaultImplWorkloadsResourceHandler) HistoryGet(namespace, name string) (*fv1.Workloads, error) {
	//return sharedK8sClient.
	//	clientSet.
	//	client.
	//	FuxiV1().
	//	Workloadses(namespace).
	//	Get(name, metav1.GetOptions{})
	return nil, nil
}

func (d *defaultImplWorkloadsResourceHandler) HistorySave(namespace string, workloads *fv1.Workloads) error {
	return nil
}

func (d *defaultImplWorkloadsResourceHandler) HistoryList(namespace string, selector labels.Selector, limit int) (
	WorkloadsSlice, error,
) {
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
		Resource(d.GetGroupVersionResource()).
		Namespace(namespace).
		List(opts)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (d *defaultImplWorkloadsResourceHandler) Get(namespace, name string) (
	runtime.Object, error,
) {
	object, err := sharedK8sClient.
		cacheInformer.
		Informer.
		ForResource(d.GetGroupVersionResource()).
		Lister().
		ByNamespace(namespace).
		Get(name)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (d *defaultImplWorkloadsResourceHandler) RemoteGet(namespace, name string, subresources ...string) (runtime.Object, error) {
	object, err := sharedK8sClient.
		cacheInformer.
		Client.
		Resource(d.GetGroupVersionResource()).
		Namespace(namespace).
		Get(name, metav1.GetOptions{}, subresources...)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (d *defaultImplWorkloadsResourceHandler) Watch(
	namespace string,
	resourceVersion string,
	timeoutSeconds int64,
	selector labels.Selector,
) (<-chan watch.Event, error) {
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
		Resource(d.GetGroupVersionResource()).
		Namespace(namespace).
		Watch(opts)
	if err != nil {
		return nil, err
	}
	return recv.ResultChan(), nil
}

func (d *defaultImplWorkloadsResourceHandler) Apply(
	namespace string,
	name string,
	obj *unstructured.Unstructured,
) (result *unstructured.Unstructured, retryErr error) {
	retryErr = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		resource := d.GetGroupVersionResource()
		getObj, getErr := sharedK8sClient.
			cacheInformer.
			Client.
			Resource(resource).
			Namespace(namespace).
			Get(name, metav1.GetOptions{})
		if errors.IsNotFound(getErr) {
			newObj, createErr := sharedK8sClient.
				cacheInformer.
				Client.
				Resource(d.GetGroupVersionResource()).
				Namespace(namespace).
				Create(obj, metav1.CreateOptions{})
			result = newObj
			return createErr
		}

		if getErr != nil {
			return getErr
		}

		if reflect.DeepEqual(getObj.Object["spec"], obj.Object["spec"]) {
			result = getObj
			return nil
		} else {
			getObj.Object["spec"] = obj.Object["spec"]
		}

		newObj, updateErr := sharedK8sClient.
			cacheInformer.
			Client.
			Resource(resource).
			Namespace(namespace).
			Update(getObj, metav1.UpdateOptions{})

		result = newObj
		return updateErr
	})

	return
}

func (d *defaultImplWorkloadsResourceHandler) Patch(namespace, name string, pathData map[string]interface{}) (*unstructured.Unstructured, error) {
	ptBytes, err := json.Marshal(pathData)
	if err != nil {
		return nil, err
	}
	return sharedK8sClient.
		cacheInformer.
		Client.
		Resource(d.GetGroupVersionResource()).
		Namespace(namespace).
		Patch(name, types.StrategicMergePatchType, ptBytes, metav1.PatchOptions{})
}

func (d *defaultImplWorkloadsResourceHandler) Delete(namespace, name string) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		return sharedK8sClient.
			cacheInformer.
			Client.
			Resource(d.GetGroupVersionResource()).
			Namespace(namespace).
			Delete(name, &metav1.DeleteOptions{})
	})
	return retryErr
}
