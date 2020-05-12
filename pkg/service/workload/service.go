package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// Service the kubernetes native resource services
type Service struct {
	WorkloadsResourceHandler
}

// NewService exported
func NewService() *Service {
	return &Service{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceService,
	}}
}
