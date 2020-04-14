package workload

// HorizontalPodAutoscaler the kubernetes native HorizontalPodAutoscaler
type HorizontalPodAutoscaler struct {
	WorkloadsResourceHandler
}

// NewHorizontalPodAutoscaler exported
func NewHorizontalPodAutoscaler() *HorizontalPodAutoscaler {
	return &HorizontalPodAutoscaler{&defaultImplWorkloadsResourceHandler{}}
}
