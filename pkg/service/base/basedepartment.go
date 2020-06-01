package base

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Base Department the kubernetes native resource deployments
type BaseDepartment struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewBaseDepartment() *BaseDepartment {
	return &BaseDepartment{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceBaseDepartment,
	}}
}
