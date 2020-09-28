package istio

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// DestinationRule Istio DestinationRule resource
type DestinationRule struct {
	common.WorkloadsResourceHandler
}

// NewDestinationRule
func NewDestinationRule() *DestinationRule {
	return &DestinationRule{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceIstioDestinationRule,
		},
	}
}
