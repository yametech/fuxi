package workload

// Deployment the kubernetes native resource deployments
type Deployment struct {
	WorkloadsResourceHandler // extended for workloadsResourceHandler
}

func NewDeployment() *Deployment {
	return &Deployment{&defaultImplWorkloadsResourceHandler{}}
}
