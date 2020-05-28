package base

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base User the kubernetes native resource deployments
type BaseUser struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBaseUser() *BaseUser {
	return &BaseUser{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceBaseUser,
	}}
}
