package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Stone is nuwa.nip.io resource Stone
type Stone struct {
	common.WorkloadsResourceHandler
}

// NewStone exported
func NewStone() *Stone {
	return &Stone{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceStone,
	}}
}
