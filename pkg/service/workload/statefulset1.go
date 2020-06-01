package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Statfulset is nuwa.nip.io resource statfulsets
type StatefulSet1 struct {
	common.WorkloadsResourceHandler
}

// NewStatfulset1 exported
func NewStatefulSet1() *StatefulSet1 {
	return &StatefulSet1{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceStatefulSet1,
	}}
}
