package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// PodSecurityPolicies the kubernetes native role
type PodSecurityPolicies struct {
	WorkloadsResourceHandler
}

// NewPodSecurityPolicies exported
func NewPodSecurityPolicies() *PodSecurityPolicies {
	return &PodSecurityPolicies{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceRole,
	}}
}
