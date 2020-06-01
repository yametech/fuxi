package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// WorkloadsTemplate is fuxi.nip.io/v1 resource workloads
type WorkloadsTemplate struct {
	common.WorkloadsResourceHandler
}

// NewWorkloadsTemplate exported
func NewWorkloadsTemplate() *WorkloadsTemplate {
	return &WorkloadsTemplate{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceWorkloadsTemplate,
		}}
}
