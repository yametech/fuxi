package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Namespace struct {
	common.WorkloadsResourceHandler
}

// NewNode exported
func NewNamespace() *Namespace {
	return &Namespace{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceNamespace}}
}
