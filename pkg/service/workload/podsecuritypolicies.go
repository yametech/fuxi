package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// PodSecurityPolicies the kubernetes native role
type PodSecurityPolicies struct {
	common.WorkloadsResourceHandler
}

// NewPodSecurityPolicies exported
func NewPodSecurityPolicies() *PodSecurityPolicies {
	return &PodSecurityPolicies{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceRole,
	}}
}
