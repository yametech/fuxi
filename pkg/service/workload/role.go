package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Role the kubernetes native role
type Role struct {
	common.WorkloadsResourceHandler
}

// NewRole exported
func NewRole() *Role {
	return &Role{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceRole,
	}}
}
