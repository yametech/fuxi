package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Ingress the kubernetes native resource ingress
type Ingress struct {
	common.WorkloadsResourceHandler
}

// NewIngress exported
func NewIngress() *Ingress {
	return &Ingress{&common.DefaultImplWorkloadsResourceHandler{dyn.ResourceIngress}}
}
