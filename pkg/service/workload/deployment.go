package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
)

// Deployment the kubernetes native resource deployments
type Deployment struct {
	WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewDeployment() *Deployment {
	return &Deployment{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceDeployment,
	}}
}
