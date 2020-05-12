package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// RoleBinding the kubernetes native role binding
type RoleBinding struct {
	WorkloadsResourceHandler
}

// NewRoleBinding exported
func NewRoleBinding() *RoleBinding {
	return &RoleBinding{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceRoleBinding,
	}}
}
