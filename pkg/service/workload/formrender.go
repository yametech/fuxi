package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

type FormRender struct {
	common.WorkloadsResourceHandler
}

// NewFormRender exported
func NewFormRender() *FormRender {
	return &FormRender{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceFormRender}}
}
