package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// StorageClass the kubernetes native resource storage class
type StorageClass struct {
	WorkloadsResourceHandler
}

// NewStorageClass exported
func NewStorageClass() *StorageClass {
	return &StorageClass{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceStorageClass,
	}}
}
