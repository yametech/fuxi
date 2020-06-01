package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Endpoint struct {
	common.WorkloadsResourceHandler
}

// NewEvent exported
func NewEndpoint() *Endpoint {
	return &Endpoint{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceEndpoint,
	}}
}
