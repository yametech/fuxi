package base

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base Permission the kubernetes native resource deployments
type BasePermission struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBasePermission() *BasePermission {
	return &BasePermission{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceBasePermission,
	}}
}
