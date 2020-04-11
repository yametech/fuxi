package workload

// CronJob is kubernetes default resource cronjob
type CronJob struct {
	WorkloadsResourceHandler // extended for workloadsResourceHandler
}

// NewCronJob exported
func NewCronJob() *CronJob {
	return &CronJob{&defaultImplWorkloadsResourceHandler{}}
}
