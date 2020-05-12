package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// Role the kubernetes native role
type Role struct {
	WorkloadsResourceHandler
}

// NewRole exported
func NewRole() *Role {
	return &Role{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceRole,
	}}
}
