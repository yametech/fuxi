package workload

// Secret the kubernetes native resource secret
type Secrets struct {
	WorkloadsResourceHandler
}

// NewSecret exported
func NewSecrets() *Secrets {
	return &Secrets{&defaultImplWorkloadsResourceHandler{}}
}
