package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

type Endpoint struct {
	WorkloadsResourceHandler
}

// NewEvent exported
func NewEndpoint() *Endpoint {
	return &Endpoint{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceEndponit,
	}}
}
