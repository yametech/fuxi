package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

type FormRender struct {
	WorkloadsResourceHandler
}

// NewFormRender exported
func NewFormRender() *FormRender {
	return &FormRender{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceFormRender,}}
}
