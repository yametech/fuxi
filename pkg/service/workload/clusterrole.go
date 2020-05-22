package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
)

// ClusterRole the kubernetes native resource ClusterRole
type ClusterRole struct {
	WorkloadsResourceHandler
}

// NewClusterRole exported
func NewClusterRole() *ClusterRole {
	return &ClusterRole{
		&defaultImplWorkloadsResourceHandler{
			dyn.ResourceClusterRole,
		}}
}
