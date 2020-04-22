package workload

type Namespace struct {
	WorkloadsResourceHandler
}

// NewNode exported
func NewNamespace() *Namespace {
	return &Namespace{&defaultImplWorkloadsResourceHandler{}}
}
