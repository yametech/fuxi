package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

type FormRender struct {
	common.WorkloadsResourceHandler
}

// NewFormRender exported
func NewFormRender() *FormRender {
	return &FormRender{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceFormRender}}
}
