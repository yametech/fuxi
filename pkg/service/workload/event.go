package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

type Event struct {
	WorkloadsResourceHandler
}

// NewEvent exported
func NewEvent() *Event {
	return &Event{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceEvent,
	}}
}
