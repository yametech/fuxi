package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// StorageClass the kubernetes native resource storage class
type StorageClass struct {
	common.WorkloadsResourceHandler
}

// NewStorageClass exported
func NewStorageClass() *StorageClass {
	return &StorageClass{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceStorageClass,
	}}
}
