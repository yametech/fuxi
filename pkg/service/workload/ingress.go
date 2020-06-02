package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Ingress the kubernetes native resource ingress
type Ingress struct {
	common.WorkloadsResourceHandler
}

// NewIngress exported
func NewIngress() *Ingress {
	return &Ingress{&common.DefaultImplWorkloadsResourceHandler{types.ResourceIngress}}
}
