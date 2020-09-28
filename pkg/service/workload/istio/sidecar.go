package istio

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Sidecar Istio Sidecar resource
type Sidecar struct {
	common.WorkloadsResourceHandler
}

// NewSidecar
func NewSidecar() *Sidecar {
	return &Sidecar{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceIstioSidecar,
		},
	}
}
