package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// HorizontalPodAutoscaler the kubernetes native HorizontalPodAutoscaler
type HorizontalPodAutoscaler struct {
	WorkloadsResourceHandler
}

// NewHorizontalPodAutoscaler exported
func NewHorizontalPodAutoscaler() *HorizontalPodAutoscaler {
	return &HorizontalPodAutoscaler{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceHorizontalPodAutoscaler}}
}
