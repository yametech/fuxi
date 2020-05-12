package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// PersistentVolume the kubernetes native resource persistent volume
type PersistentVolume struct {
	WorkloadsResourceHandler
}

// NewPersistentVolume exported
func NewPersistentVolume() *PersistentVolume {
	return &PersistentVolume{&defaultImplWorkloadsResourceHandler{
		dyn.ResourcePersistentVolume,
	}}
}
