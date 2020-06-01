package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Form the kubernetes native service account
type Form struct {
	common.WorkloadsResourceHandler
}

// NewField exported
func NewForm() *Form {
	return &Form{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceForm,
	}}
}
