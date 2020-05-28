package base

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base Role the kubernetes native resource deployments
type BaseRole struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBaseRole() *BaseRole {
	return &BaseRole{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceBaseRole,
	}}
}
