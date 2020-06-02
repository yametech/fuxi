package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// NetworkPolicy the kubernetes native resource network policy
type NetworkPolicy struct {
	common.WorkloadsResourceHandler
}

// NewNetworkPolicy exported
func NewNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceNetworkPolicy,
	}}
}
