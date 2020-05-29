package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
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
			dyn.ResourcePieline,
		},
	}
}
