package base

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base RoleUser the kubernetes native resource deployments
type BaseRoleUser struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBaseRoleUser() *BaseRoleUser {
	return &BaseRoleUser{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceBaseRoleUser,
	}}
}
