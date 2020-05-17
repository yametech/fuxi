package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// Statfulset is nuwa.nip.io resource statfulsets
type StatefulSet1 struct {
	WorkloadsResourceHandler
}

// NewStatfulset1 exported
func NewStatefulSet1() *StatefulSet1 {
	return &StatefulSet1{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceStatefulSet1,
	}}
}
