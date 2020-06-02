package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Statfulset is kubernetes default resource statfulsets
type StatefulSet struct {
	common.WorkloadsResourceHandler
}

// NewStatfulset exported
func NewStatefulSet() *StatefulSet {
	return &StatefulSet{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceStatefulSet,
	}}
}
