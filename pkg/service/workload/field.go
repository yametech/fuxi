package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Field the kubernetes native service account
type Field struct {
	common.WorkloadsResourceHandler
}

// NewField exported
func NewField() *Field {
	return &Field{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceField,
	}}
}
