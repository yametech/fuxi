package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Event struct {
	common.WorkloadsResourceHandler
}

// NewEvent exported
func NewEvent() *Event {
	return &Event{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceEvent,
	}}
}
