package workload

// DaemonSet the kubernetes native resource daemonsets
type DaemonSet struct {
	WorkloadsResourceHandler
}

// NewDaemonSet exported
func NewDaemonSet() *DaemonSet {
	return &DaemonSet{&defaultImplWorkloadsResourceHandler{}}
}
