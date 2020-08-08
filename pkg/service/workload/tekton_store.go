package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// TektonStore is kubernetes default resource tekton graph
type TektonStore struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

// NewTektonStore exported
func NewTektonStore() *TektonStore {
	return &TektonStore{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceTektonStore,
		},
	}
}
