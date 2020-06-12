package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

type SubNet struct {
	common.WorkloadsResourceHandler
}

// NewSubNet exported
func NewSubNet() *SubNet {
	return &SubNet{&common.DefaultImplWorkloadsResourceHandler{types.ResourceSubNet}}
}
