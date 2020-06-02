package base

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base Role the kubernetes native resource deployments
type BaseRole struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBaseRole() *BaseRole {
	return &BaseRole{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceBaseRole,
	}}
}
