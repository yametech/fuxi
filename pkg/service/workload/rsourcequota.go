package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// ResourceQuota the kubernetes native resource resource quota
type ResourceQuota struct {
	WorkloadsResourceHandler
}

// NewResourceQuota exported
func NewResourceQuota() *ResourceQuota {
	return &ResourceQuota{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceResourceQuota,
	}}
}
