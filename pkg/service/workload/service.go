package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Service the kubernetes native resource services
type Service struct {
	common.WorkloadsResourceHandler
}

// NewService exported
func NewService() *Service {
	return &Service{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceService,
	}}
}
