package workload

// Statfulset is kubernetes default resource statfulsets
type StatefulSet struct {
	WorkloadsResourceHandler
}

// NewStatfulset exported
func NewStatefulSet() *StatefulSet {
	return &StatefulSet{&defaultImplWorkloadsResourceHandler{}}
}
