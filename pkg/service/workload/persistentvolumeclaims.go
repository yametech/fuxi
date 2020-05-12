package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// PersistentVolumeClaims the kubernetes native resource persistent volume claims
type PersistentVolumeClaims struct {
	WorkloadsResourceHandler
}

// NewPersistentVolumeClaims exported
func NewPersistentVolumeClaims() *PersistentVolumeClaims {
	return &PersistentVolumeClaims{&defaultImplWorkloadsResourceHandler{
		dyn.ResourcePersistentVolumeClaims,
	}}
}
