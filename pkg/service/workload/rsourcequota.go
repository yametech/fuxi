package workload

// ResourceQuota the kubernetes native resource resource quota
type ResourceQuota struct {
	WorkloadsResourceHandler
}

// NewResourceQuota exported
func NewResourceQuota() *ResourceQuota {
	return &ResourceQuota{&defaultImplWorkloadsResourceHandler{}}
}
