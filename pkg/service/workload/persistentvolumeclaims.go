package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// PersistentVolumeClaims the kubernetes native resource persistent volume claims
type PersistentVolumeClaims struct {
	common.WorkloadsResourceHandler
}

// NewPersistentVolumeClaims exported
func NewPersistentVolumeClaims() *PersistentVolumeClaims {
	return &PersistentVolumeClaims{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourcePersistentVolumeClaims,
	}}
}
