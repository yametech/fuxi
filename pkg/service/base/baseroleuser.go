package base

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base RoleUser the kubernetes native resource deployments
type BaseRoleUser struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBaseRoleUser() *BaseRoleUser {
	return &BaseRoleUser{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceBaseRoleUser,
	}}
}
