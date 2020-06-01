package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Namespace struct {
	common.WorkloadsResourceHandler
}

// NewNode exported
func NewNamespace() *Namespace {
	return &Namespace{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceNamespace}}
}
