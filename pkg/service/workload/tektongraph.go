package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// TektonGraph is kubernetes default resource tekton graph
type TektonGraph struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

// NewTektonGraph exported
func NewTektonGraph() *TektonGraph {
	return &TektonGraph{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceTektonGraph,
		},
	}
}
