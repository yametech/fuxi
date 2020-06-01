package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
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
			types.ResourceCustomResourceDefinition,
		}}
}
