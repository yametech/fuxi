package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// NetworkPolicy the kubernetes native resource network policy
type NetworkPolicy struct {
	common.WorkloadsResourceHandler
}

// NewNetworkPolicy exported
func NewNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceNetworkPolicy,
	}}
}
