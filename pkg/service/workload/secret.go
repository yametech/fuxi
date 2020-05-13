package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
)

// Secret the kubernetes native resource secret
type Secrets struct {
	WorkloadsResourceHandler
}

// NewSecret exported
func NewSecrets() *Secrets {
	return &Secrets{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceSecrets,
	}}
}
