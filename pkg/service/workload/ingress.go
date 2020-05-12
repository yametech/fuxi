package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// Ingress the kubernetes native resource ingress
type Ingress struct {
	WorkloadsResourceHandler
}

// NewIngress exported
func NewIngress() *Ingress {
	return &Ingress{&defaultImplWorkloadsResourceHandler{ dyn.ResourceIngress}}
}
