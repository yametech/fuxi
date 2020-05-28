package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
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
			dyn.ResourceClusterRole,
		}}
}
