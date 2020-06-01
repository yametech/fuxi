package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Deployment the kubernetes native resource deployments
type Deployment struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewDeployment() *Deployment {
	return &Deployment{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceDeployment,
	}}
}
