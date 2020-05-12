package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// CustomResourceDefinition the kubernetes native CustomResourceDefinition
type CustomResourceDefinition struct {
	WorkloadsResourceHandler
}

// NewCustomResourceDefinition exported
func NewCustomResourceDefinition() *CustomResourceDefinition {
	return &CustomResourceDefinition{
		&defaultImplWorkloadsResourceHandler{
			dyn.ResourceCustomResourceDefinition,
		}}
}
