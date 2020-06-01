package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// TaskRun the kubernetes Knative tekton taskrun resource
type TaskRun struct {
	common.WorkloadsResourceHandler
}

// NewTaskRun exported
func NewTaskRun() *TaskRun {
	return &TaskRun{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceTaskRun,
		},
	}
}
