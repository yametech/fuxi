package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
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
			types.ResourceClusterRoleBinding,
		}}
}
