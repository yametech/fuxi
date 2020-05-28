package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ServiceAccount the kubernetes native service account
type ServiceAccount struct {
	common.WorkloadsResourceHandler
}

// NewServiceAccount exported
func NewServiceAccount() *ServiceAccount {
	return &ServiceAccount{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceServiceAccount,
	}}
}
