package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// Service the kubernetes native resource services
type Service struct {
	common.WorkloadsResourceHandler
}

// NewService exported
func NewService() *Service {
	return &Service{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceService,
	}}
}
