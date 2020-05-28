package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Form the kubernetes native service account
type Form struct {
	common.WorkloadsResourceHandler
}

// NewField exported
func NewForm() *Form {
	return &Form{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceForm,
	}}
}
