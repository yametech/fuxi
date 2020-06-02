package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// HorizontalPodAutoscaler the kubernetes native HorizontalPodAutoscaler
type HorizontalPodAutoscaler struct {
	common.WorkloadsResourceHandler
}

// NewHorizontalPodAutoscaler exported
func NewHorizontalPodAutoscaler() *HorizontalPodAutoscaler {
	return &HorizontalPodAutoscaler{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceHorizontalPodAutoscaler}}
}
