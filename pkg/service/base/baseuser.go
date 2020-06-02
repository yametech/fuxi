package base

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base User the kubernetes native resource deployments
type BaseUser struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBaseUser() *BaseUser {
	return &BaseUser{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceBaseUser,
	}}
}
