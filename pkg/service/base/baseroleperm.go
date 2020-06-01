package base

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base RolePerm the kubernetes native resource deployments
type BaseRolePerm struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBaseRolePerm() *BaseRolePerm {
	return &BaseRolePerm{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceBaseRolePerm,
	}}
}
