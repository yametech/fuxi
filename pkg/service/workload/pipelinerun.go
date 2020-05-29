package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// PipelineRun the kubernetes Knative tekton pipeline run resource
type PipelineRun struct {
	common.WorkloadsResourceHandler
}

// NewPipelineRun exported
func NewPipelineRun() *PipelineRun {
	return &PipelineRun{
		&common.DefaultImplWorkloadsResourceHandler{
			dyn.ResourcePipelineRun,
		},
	}
}
