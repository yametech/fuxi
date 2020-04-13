package workload

// NetworkPolicy the kubernetes native resource network policy
type NetworkPolicy struct {
	WorkloadsResourceHandler
}

// NewNetworkPolicy exported
func NewNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{&defaultImplWorkloadsResourceHandler{}}
}
