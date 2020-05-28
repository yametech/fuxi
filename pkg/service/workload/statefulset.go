package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Statfulset is kubernetes default resource statfulsets
type StatefulSet struct {
	common.WorkloadsResourceHandler
}

// NewStatfulset exported
func NewStatefulSet() *StatefulSet {
	return &StatefulSet{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceStatefulSet,
	}}
}
