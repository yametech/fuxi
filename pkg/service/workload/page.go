package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Page the kubernetes native service account
type Page struct {
	common.WorkloadsResourceHandler
}

// NewPage exported
func NewPage() *Page {
	return &Page{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourcePage,
	}}
}
