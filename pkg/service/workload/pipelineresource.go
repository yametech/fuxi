package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// PipelineResource the kubernetes Knative tekton PipelineResource resource
type PipelineResource struct {
	common.WorkloadsResourceHandler
}

// NewPipelineResource exported
func NewPipelineResource() *PipelineResource {
	return &PipelineResource{
		&common.DefaultImplWorkloadsResourceHandler{
			dyn.ResourcePipelineResource,
		},
	}
}
