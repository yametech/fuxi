package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ClusterRole the kubernetes native resource ClusterRole
type ClusterRole struct {
	common.WorkloadsResourceHandler
}

// NewClusterRole exported
func NewClusterRole() *ClusterRole {
	return &ClusterRole{
		&common.DefaultImplWorkloadsResourceHandler{
			types.ResourceClusterRole,
		}}
}
