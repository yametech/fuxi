package workload

type Generic struct {
	WorkloadsResourceHandler
}

func NewGeneric() *Generic {
	return &Generic{&defaultImplWorkloadsResourceHandler{}}
}
