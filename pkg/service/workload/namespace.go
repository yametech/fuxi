package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
)

type Namespace struct {
	WorkloadsResourceHandler
}

// NewNode exported
func NewNamespace() *Namespace {
	return &Namespace{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceNamespace}}
}
