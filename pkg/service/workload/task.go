package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Task the kubernetes Knative tekton Task resource
type Task struct {
	common.WorkloadsResourceHandler
}

// NewTask exported
func NewTask() *Task {
	return &Task{
		&common.DefaultImplWorkloadsResourceHandler{
			dyn.ResourceTask,
		},
	}
}
