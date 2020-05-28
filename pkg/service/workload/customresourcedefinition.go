package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// CustomResourceDefinition the kubernetes native CustomResourceDefinition
type CustomResourceDefinition struct {
	common.WorkloadsResourceHandler
}

// NewCustomResourceDefinition exported
func NewCustomResourceDefinition() *CustomResourceDefinition {
	return &CustomResourceDefinition{
		&common.DefaultImplWorkloadsResourceHandler{
			dyn.ResourceCustomResourceDefinition,
		}}
}
