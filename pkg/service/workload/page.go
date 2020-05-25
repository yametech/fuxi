package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// Page the kubernetes native service account
type Page struct {
	WorkloadsResourceHandler
}

// NewPage exported
func NewPage() *Page {
	return &Page{&defaultImplWorkloadsResourceHandler{
		dyn.ResourcePage,
	}}
}
