package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Endpoint struct {
	common.WorkloadsResourceHandler
}

// NewEvent exported
func NewEndpoint() *Endpoint {
	return &Endpoint{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceEndponit,
	}}
}
