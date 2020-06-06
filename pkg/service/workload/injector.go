package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Injector is nuwa.nip.io resource Injector
type Injector struct {
	common.WorkloadsResourceHandler
}

// NewInjector exported
func NewInjector() *Injector {
	return &Injector{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceInjector,
	}}
}
