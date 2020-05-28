package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Event struct {
	common.WorkloadsResourceHandler
}

// NewEvent exported
func NewEvent() *Event {
	return &Event{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceEvent,
	}}
}
