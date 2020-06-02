package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Pipeline the kubernetes Knative tekton pipeline resource
type Pipeline struct {
	common.WorkloadsResourceHandler
}

// NewPipeline exported
func NewPipeline() *Pipeline {
	return &Pipeline{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourcePieline,
		},
	}
}
