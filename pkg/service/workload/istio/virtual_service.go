package istio

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// VirtualService Istio VirtualService resource
type VirtualService struct {
	common.WorkloadsResourceHandler
}

// NewVirtualService
func NewVirtualService() *VirtualService {
	return &VirtualService{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceIstioSidecar,
		},
	}
}
