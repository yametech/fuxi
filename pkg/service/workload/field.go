package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// Field the kubernetes native service account
type Field struct {
	WorkloadsResourceHandler
}

// NewField exported
func NewField() *Field {
	return &Field{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceField,
	}}
}
