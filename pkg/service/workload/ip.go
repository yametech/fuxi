package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

type IP struct {
	common.WorkloadsResourceHandler
}

// NewIP exported
func NewIP() *IP {
	return &IP{&common.DefaultImplWorkloadsResourceHandler{types.ResourceIP}}
}
