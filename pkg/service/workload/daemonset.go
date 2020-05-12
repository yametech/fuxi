package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
)

// DaemonSet the kubernetes native resource daemonsets
type DaemonSet struct {
	WorkloadsResourceHandler
}

// NewDaemonSet exported
func NewDaemonSet() *DaemonSet {
	return &DaemonSet{&defaultImplWorkloadsResourceHandler{dyn.ResourceDaemonSet}}
}
