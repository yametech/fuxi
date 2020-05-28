package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// PersistentVolume the kubernetes native resource persistent volume
type PersistentVolume struct {
	common.WorkloadsResourceHandler
}

// NewPersistentVolume exported
func NewPersistentVolume() *PersistentVolume {
	return &PersistentVolume{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourcePersistentVolume,
	}}
}
