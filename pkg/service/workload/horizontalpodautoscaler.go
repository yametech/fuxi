package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// HorizontalPodAutoscaler the kubernetes native HorizontalPodAutoscaler
type HorizontalPodAutoscaler struct {
	common.WorkloadsResourceHandler
}

// NewHorizontalPodAutoscaler exported
func NewHorizontalPodAutoscaler() *HorizontalPodAutoscaler {
	return &HorizontalPodAutoscaler{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceHorizontalPodAutoscaler}}
}
