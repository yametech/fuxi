package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Secret the kubernetes native resource secret
type Secrets struct {
	WorkloadsResourceHandler
}

func (r *Secrets) GetGroupVersionResource() schema.GroupVersionResource {
	return dyn.ResourceSecrets
}

// NewSecret exported
func NewSecrets() *Secrets {
	return &Secrets{&defaultImplWorkloadsResourceHandler{}}
}
