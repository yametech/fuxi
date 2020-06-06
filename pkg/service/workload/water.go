package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Water is nuwa.nip.io resource Water
type Water struct {
	common.WorkloadsResourceHandler
}

// NewWater exported
func NewWater() *Water {
	return &Water{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceWater,
	}}
}
