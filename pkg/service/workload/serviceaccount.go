package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// ServiceAccount the kubernetes native service account
type ServiceAccount struct {
	WorkloadsResourceHandler
}

// NewServiceAccount exported
func NewServiceAccount() *ServiceAccount {
	return &ServiceAccount{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceServiceAccount,
	}}
}
