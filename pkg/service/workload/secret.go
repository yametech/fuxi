package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Secret the kubernetes native resource secret
type Secrets struct {
	common.WorkloadsResourceHandler
}

// NewSecret exported
func NewSecrets() *Secrets {
	return &Secrets{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceSecrets,
	}}
}
