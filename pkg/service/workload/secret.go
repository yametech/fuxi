package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Secret the kubernetes native resource secret
type Secrets struct {
	common.WorkloadsResourceHandler
}

// NewSecret exported
func NewSecrets() *Secrets {
	return &Secrets{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceSecrets,
	}}
}
