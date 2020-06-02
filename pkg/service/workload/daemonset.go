package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// DaemonSet the kubernetes native resource daemonsets
type DaemonSet struct {
	common.WorkloadsResourceHandler
}

// NewDaemonSet exported
func NewDaemonSet() *DaemonSet {
	return &DaemonSet{&common.DefaultImplWorkloadsResourceHandler{types.ResourceDaemonSet}}
}
