package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// RoleBinding the kubernetes native role binding
type RoleBinding struct {
	common.WorkloadsResourceHandler
}

// NewRoleBinding exported
func NewRoleBinding() *RoleBinding {
	return &RoleBinding{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceRoleBinding,
	}}
}
