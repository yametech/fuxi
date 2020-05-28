package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ClusterRoleBinding the kubernetes native resource configmaps
type ClusterRoleBinding struct {
	common.WorkloadsResourceHandler
}

// NewClusterRoleBinding exported
func NewClusterRoleBinding() *ClusterRoleBinding {
	return &ClusterRoleBinding{
		&common.DefaultImplWorkloadsResourceHandler{
			dyn.ResourceClusterRoleBinding,
		}}
}
