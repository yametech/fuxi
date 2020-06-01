package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ResourceQuota the kubernetes native resource resource quota
type ResourceQuota struct {
	common.WorkloadsResourceHandler
}

// NewResourceQuota exported
func NewResourceQuota() *ResourceQuota {
	return &ResourceQuota{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceResourceQuota,
	}}
}
