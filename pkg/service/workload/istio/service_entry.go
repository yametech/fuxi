package istio

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ServiceEntry Istio ServiceEntry resource
type ServiceEntry struct {
	common.WorkloadsResourceHandler
}

// NewServiceEntry
func NewServiceEntry() *ServiceEntry {
	return &ServiceEntry{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceIstioServiceEntry,
		},
	}
}
