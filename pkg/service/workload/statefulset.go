package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// Statfulset is kubernetes default resource statfulsets
type StatefulSet struct {
	WorkloadsResourceHandler
}

// NewStatfulset exported
func NewStatefulSet() *StatefulSet {
	return &StatefulSet{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceStatefulSet,
	}}
}
