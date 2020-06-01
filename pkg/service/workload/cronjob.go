package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// CronJob is kubernetes default resource cronjob
type CronJob struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

// NewCronJob exported
func NewCronJob() *CronJob {
	return &CronJob{&common.DefaultImplWorkloadsResourceHandler{types.ResourceCronJobs}}
}
