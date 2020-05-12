package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

type Job struct {
	WorkloadsResourceHandler
}

// NewJob exported
func NewJob() *Job {
	return &Job{&defaultImplWorkloadsResourceHandler{dyn.ResourceJob}}
}
