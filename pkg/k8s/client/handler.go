package client

import (
	"fmt"

	"github.com/yametech/fuxi/pkg/k8s/client/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

type ResourceHandler struct {
	client       *kubernetes.Clientset
	cacheFactory *CacheFactory
}

func NewResourceHandler(kubeClient *kubernetes.Clientset, cacheFactory *CacheFactory) *ResourceHandler {
	return &ResourceHandler{
		client:       kubeClient,
		cacheFactory: cacheFactory,
	}
}

func (h *ResourceHandler) Create(kind string, namespace string, object *runtime.Unknown) (*runtime.Unknown, error) {
	resource, ok := api.KindToResourceMap[kind]
	if !ok {
		return nil, fmt.Errorf("Resource kind (%s) not support yet . ", kind)
	}
	kubeClient := h.getClientByGroupVersion(resource.GroupVersionResourceKind.GroupVersionResource)
	req := kubeClient.Post().
		Resource(kind).
		SetHeader("Content-Type", "application/json").
		Body([]byte(object.Raw))
	if resource.Namespaced {
		req.Namespace(namespace)
	}
	var result runtime.Unknown
	err := req.Do().Into(&result)

	return &result, err
}

func (h *ResourceHandler) Update(kind string, namespace string, name string, object *runtime.Unknown) (*runtime.Unknown, error) {
	resource, ok := api.KindToResourceMap[kind]
	if !ok {
		return nil, fmt.Errorf("Resource kind (%s) not support yet . ", kind)
	}

	kubeClient := h.getClientByGroupVersion(resource.GroupVersionResourceKind.GroupVersionResource)
	req := kubeClient.Put().
		Resource(kind).
		Name(name).
		SetHeader("Content-Type", "application/json").
		Body([]byte(object.Raw))
	if resource.Namespaced {
		req.Namespace(namespace)
	}

	var result runtime.Unknown
	err := req.Do().Into(&result)

	return &result, err
}

func (h *ResourceHandler) Delete(kind string, namespace string, name string, options *metav1.DeleteOptions) error {
	resource, ok := api.KindToResourceMap[kind]
	if !ok {
		return fmt.Errorf("Resource kind (%s) not support yet . ", kind)
	}
	kubeClient := h.getClientByGroupVersion(resource.GroupVersionResourceKind.GroupVersionResource)
	req := kubeClient.Delete().
		Resource(kind).
		Name(name).
		Body(options)
	if resource.Namespaced {
		req.Namespace(namespace)
	}

	return req.Do().Error()
}

// Get object from cache
func (h *ResourceHandler) Get(kind string, namespace string, name string) (runtime.Object, error) {
	resource, ok := api.KindToResourceMap[kind]
	if !ok {
		return nil, fmt.Errorf("Resource kind (%s) not support yet . ", kind)
	}
	genericInformer, err := h.cacheFactory.sharedInformerFactory.ForResource(resource.GroupVersionResourceKind.GroupVersionResource)
	if err != nil {
		return nil, err
	}
	lister := genericInformer.Lister()
	var result runtime.Object
	if resource.Namespaced {
		result, err = lister.ByNamespace(namespace).Get(name)
		if err != nil {
			return nil, err
		}
	} else {
		result, err = lister.Get(name)
		if err != nil {
			return nil, err
		}
	}
	result.GetObjectKind().SetGroupVersionKind(
		schema.GroupVersionKind{
			Group:   resource.GroupVersionResourceKind.Group,
			Version: resource.GroupVersionResourceKind.Version,
			Kind:    resource.GroupVersionResourceKind.Kind,
		})

	return result, nil
}

// Get object from cache
func (h *ResourceHandler) List(kind string, namespace string, labelSelector string) ([]runtime.Object, error) {
	resource, ok := api.KindToResourceMap[kind]
	if !ok {
		return nil, fmt.Errorf("Resource kind (%s) not support yet . ", kind)
	}
	genericInformer, err := h.cacheFactory.sharedInformerFactory.ForResource(resource.GroupVersionResourceKind.GroupVersionResource)
	if err != nil {
		return nil, err
	}
	selectors, err := labels.Parse(labelSelector)
	if err != nil {
		return nil, err
	}

	lister := genericInformer.Lister()
	var objs []runtime.Object
	if resource.Namespaced {
		objs, err = lister.ByNamespace(namespace).List(selectors)
		if err != nil {
			return nil, err
		}
	} else {
		objs, err = lister.List(selectors)
		if err != nil {
			return nil, err
		}
	}

	for i := range objs {
		objs[i].GetObjectKind().SetGroupVersionKind(
			schema.GroupVersionKind{
				Group:   resource.GroupVersionResourceKind.Group,
				Version: resource.GroupVersionResourceKind.Version,
				Kind:    resource.GroupVersionResourceKind.Kind,
			})
	}

	return objs, nil
}
