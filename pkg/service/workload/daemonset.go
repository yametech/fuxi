package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// DaemonSet the kubernetes native resource daemonsets
type DaemonSet struct {
	common.WorkloadsResourceHandler
}

// NewDaemonSet exported
func NewDaemonSet() *DaemonSet {
	return &DaemonSet{&common.DefaultImplWorkloadsResourceHandler{dyn.ResourceDaemonSet}}
}
