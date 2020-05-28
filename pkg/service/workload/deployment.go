package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Deployment the kubernetes native resource deployments
type Deployment struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewDeployment() *Deployment {
	return &Deployment{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceDeployment,
	}}
}
