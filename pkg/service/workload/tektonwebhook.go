package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// TektonWebHook is kubernetes default resource tekton graph
type TektonWebHook struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

// NewTektonWebHook exported
func NewTektonWebHook() *TektonWebHook {
	return &TektonWebHook{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceTektonWebHook,
		},
	}
}
