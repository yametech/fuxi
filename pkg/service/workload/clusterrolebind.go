package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
)

// ClusterRoleBinding the kubernetes native resource configmaps
type ClusterRoleBinding struct {
	WorkloadsResourceHandler
}

// NewClusterRoleBinding exported
func NewClusterRoleBinding() *ClusterRoleBinding {
	return &ClusterRoleBinding{
		&defaultImplWorkloadsResourceHandler{
			dyn.ResourceClusterRoleBinding,
		}}
}
