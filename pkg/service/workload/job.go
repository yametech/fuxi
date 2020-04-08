package workload

type Job struct {
	WorkloadsResourceHandler
}

// NewJob exported
func NewJob() *Job {
	return &Job{&defaultImplWorkloadsResourceHandler{}}
}
