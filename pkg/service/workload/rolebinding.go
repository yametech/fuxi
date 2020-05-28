package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// RoleBinding the kubernetes native role binding
type RoleBinding struct {
	common.WorkloadsResourceHandler
}

// NewRoleBinding exported
func NewRoleBinding() *RoleBinding {
	return &RoleBinding{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceRoleBinding,
	}}
}
