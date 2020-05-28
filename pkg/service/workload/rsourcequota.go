package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ResourceQuota the kubernetes native resource resource quota
type ResourceQuota struct {
	common.WorkloadsResourceHandler
}

// NewResourceQuota exported
func NewResourceQuota() *ResourceQuota {
	return &ResourceQuota{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceResourceQuota,
	}}
}
