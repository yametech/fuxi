package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// WorkloadsTemplate is fuxi.nip.io/v1 resource workloads
type WorkloadsTemplate struct {
	WorkloadsResourceHandler
}

// NewWorkloadsTemplate exported
func NewWorkloadsTemplate() *WorkloadsTemplate {
	return &WorkloadsTemplate{
		&defaultImplWorkloadsResourceHandler{
			dyn.ResourceWorkloadsTemplate,
		}}
}
