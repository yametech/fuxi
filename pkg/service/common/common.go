package common

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
	//"sort"
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

type WorkloadsResourceVersion interface {
	SetGroupVersionResource(schema.GroupVersionResource)
	GetGroupVersionResource() schema.GroupVersionResource
}

// WorkloadsResourceHandler all needed interface defined
type WorkloadsResourceHandler interface {
	ResourceQuery
	ResourceApply
	WorkloadsResourceVersion
}

// check the default implemented
var _ WorkloadsResourceHandler = &DefaultImplWorkloadsResourceHandler{}

type DefaultImplWorkloadsResourceHandler struct {
	GroupVersionResource schema.GroupVersionResource
}

func (d *DefaultImplWorkloadsResourceHandler) GetGroupVersionResource() schema.GroupVersionResource {
	return d.GroupVersionResource
}
func (d *DefaultImplWorkloadsResourceHandler) SetGroupVersionResource(g schema.GroupVersionResource) {
	d.GroupVersionResource = g
}
func (d *DefaultImplWorkloadsResourceHandler) List(
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
	items, err = SharedK8sClient.
		ClientV2.
		Interface.
		Resource(d.GetGroupVersionResource()).
		Namespace(namespace).
		List(opts)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (d *DefaultImplWorkloadsResourceHandler) Get(namespace, name string) (
	runtime.Object, error,
) {
	object, err := SharedK8sClient.
		ClientV2.
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

func (d *DefaultImplWorkloadsResourceHandler) RemoteGet(namespace, name string, subresources ...string) (runtime.Object, error) {
	object, err := SharedK8sClient.
		ClientV2.
		Interface.
		Resource(d.GetGroupVersionResource()).
		Namespace(namespace).
		Get(name, metav1.GetOptions{}, subresources...)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (d *DefaultImplWorkloadsResourceHandler) Watch(
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
	recv, err := SharedK8sClient.
		ClientV2.
		Interface.
		Resource(d.GetGroupVersionResource()).
		Namespace(namespace).
		Watch(opts)
	if err != nil {
		return nil, err
	}
	return recv.ResultChan(), nil
}

func (d *DefaultImplWorkloadsResourceHandler) Apply(
	namespace string,
	name string,
	obj *unstructured.Unstructured,
) (result *unstructured.Unstructured, retryErr error) {
	retryErr = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		resource := d.GetGroupVersionResource()
		getObj, getErr := SharedK8sClient.
			ClientV2.
			Interface.
			Resource(resource).
			Namespace(namespace).
			Get(name, metav1.GetOptions{})
		if errors.IsNotFound(getErr) {
			newObj, createErr := SharedK8sClient.
				ClientV2.
				Interface.
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

		newObj, updateErr := SharedK8sClient.
			ClientV2.
			Interface.
			Resource(resource).
			Namespace(namespace).
			Update(getObj, metav1.UpdateOptions{})

		result = newObj
		return updateErr
	})

	return
}

func (d *DefaultImplWorkloadsResourceHandler) Patch(namespace, name string, pathData map[string]interface{}) (*unstructured.Unstructured, error) {
	ptBytes, err := json.Marshal(pathData)
	if err != nil {
		return nil, err
	}
	return SharedK8sClient.
		ClientV2.
		Interface.
		Resource(d.GetGroupVersionResource()).
		Namespace(namespace).
		Patch(name, types.StrategicMergePatchType, ptBytes, metav1.PatchOptions{})
}

func (d *DefaultImplWorkloadsResourceHandler) Delete(namespace, name string) error {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		return SharedK8sClient.
			ClientV2.
			Interface.
			Resource(d.GetGroupVersionResource()).
			Namespace(namespace).
			Delete(name, &metav1.DeleteOptions{})
	})
	return retryErr
}
