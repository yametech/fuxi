package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Role the kubernetes native role
type Role struct {
	common.WorkloadsResourceHandler
}

// NewRole exported
func NewRole() *Role {
	return &Role{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceRole,
	}}
}
