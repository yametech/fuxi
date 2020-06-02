package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Job struct {
	common.WorkloadsResourceHandler
}

// NewJob exported
func NewJob() *Job {
	return &Job{&common.DefaultImplWorkloadsResourceHandler{types.ResourceJob}}
}
