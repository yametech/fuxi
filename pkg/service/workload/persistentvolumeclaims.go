package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// PersistentVolumeClaims the kubernetes native resource persistent volume claims
type PersistentVolumeClaims struct {
	common.WorkloadsResourceHandler
}

// NewPersistentVolumeClaims exported
func NewPersistentVolumeClaims() *PersistentVolumeClaims {
	return &PersistentVolumeClaims{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourcePersistentVolumeClaims,
	}}
}
