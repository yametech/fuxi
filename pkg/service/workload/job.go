package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Job struct {
	common.WorkloadsResourceHandler
}

// NewJob exported
func NewJob() *Job {
	return &Job{&common.DefaultImplWorkloadsResourceHandler{dyn.ResourceJob}}
}
