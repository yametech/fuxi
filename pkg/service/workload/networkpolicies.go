package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// NetworkPolicy the kubernetes native resource network policy
type NetworkPolicy struct {
	WorkloadsResourceHandler
}

// NewNetworkPolicy exported
func NewNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceNetworkPolicy,
	}}
}
