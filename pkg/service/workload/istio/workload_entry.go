package istio

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// WorkloadEntry Istio WorkloadEntry resource
type WorkloadEntry struct {
	common.WorkloadsResourceHandler
}

// NewWorkloadEntry
func NewWorkloadEntry() *WorkloadEntry {
	return &WorkloadEntry{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceIstioServiceEntry,
		},
	}
}
