package base

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base Permission the kubernetes native resource deployments
type BasePermission struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBasePermission() *BasePermission {
	return &BasePermission{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceBasePermission,
	}}
}
