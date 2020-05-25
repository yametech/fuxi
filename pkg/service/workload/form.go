package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// Form the kubernetes native service account
type Form struct {
	WorkloadsResourceHandler
}

// NewField exported
func NewForm() *Form {
	return &Form{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceForm,
	}}
}
