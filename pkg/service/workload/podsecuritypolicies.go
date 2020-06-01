package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// PodSecurityPolicies the kubernetes native role
type PodSecurityPolicies struct {
	common.WorkloadsResourceHandler
}

// NewPodSecurityPolicies exported
func NewPodSecurityPolicies() *PodSecurityPolicies {
	return &PodSecurityPolicies{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceRole,
	}}
}
