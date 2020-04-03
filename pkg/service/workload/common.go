package workload

import (
	fv1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	LableForResourceTypeHistory = "history" // "history":"${RESOURCE_NAME}"
	LableIsLatest               = "latest"  // "latest":true
)

type WorkloadsSlice []*fv1.Workloads

func (w WorkloadsSlice) Len() int      { return len(w) }
func (w WorkloadsSlice) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
func (w WorkloadsSlice) Less(i, j int) bool {
	if w[i].ObjectMeta.CreationTimestamp.Before(&w[j].ObjectMeta.CreationTimestamp) {
		return true
	}
	return false
}

type ResourceQuery interface {
	List(res, pos int64, size int64, flag string) (list *unstructured.UnstructuredList, err error)
	Get(res, name string, resourceVersion string) (item *unstructured.Unstructured, err error)
	Watch(res, name string) (itemChan chan *unstructured.Unstructured, closed chan struct{}, err error)
}

type ResourceApply interface {
	Apply(res, obj *unstructured.Unstructured) error
	Delete(res, namespace, name string) error
}

type HistoryQuery interface {
	HistoryList(res, namespace, name string, limit int) (WorkloadsSlice, error)
}

type WorkloadsResourceHandler interface {
	HistoryQuery
	ResourceQuery
	ResourceApply
}

var _ WorkloadsResourceHandler = &defaultImplWorkloadsResourceHandler{}

type defaultImplWorkloadsResourceHandler struct{}

func (d defaultImplWorkloadsResourceHandler) HistoryList(res, namespace, name string, limit int) (WorkloadsSlice, error) {
	panic("implement me")
}

func (d defaultImplWorkloadsResourceHandler) List(res, pos int64, size int64, flag string) (list *unstructured.UnstructuredList, err error) {
	panic("implement me")
}

func (d defaultImplWorkloadsResourceHandler) Get(res, name string, resourceVersion string) (item *unstructured.Unstructured, err error) {
	panic("implement me")
}

func (d defaultImplWorkloadsResourceHandler) Watch(res, name string) (itemChan chan *unstructured.Unstructured, closed chan struct{}, err error) {
	panic("implement me")
}

func (d defaultImplWorkloadsResourceHandler) Apply(res, obj *unstructured.Unstructured) error {
	panic("implement me")
}

func (d defaultImplWorkloadsResourceHandler) Delete(res, namespace, name string) error {
	panic("implement me")
}

