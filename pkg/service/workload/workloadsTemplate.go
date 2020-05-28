package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
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
			dyn.ResourceWorkloadsTemplate,
		}}
}
