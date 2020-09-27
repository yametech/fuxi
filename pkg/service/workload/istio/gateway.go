package istio

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Gateway Istio Gateway resource
type Gateway struct {
	common.WorkloadsResourceHandler
}

// NewGateway
func NewGateway() *Gateway {
	return &Gateway{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceIstioGateway,
		},
	}
}
