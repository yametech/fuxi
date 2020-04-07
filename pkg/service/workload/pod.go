package workload

// Pod doc kubernetes
type Pod struct {
	WorkloadsResourceHandler
}

func NewPod() *Pod {
	return &Pod{&defaultImplWorkloadsResourceHandler{}}
}
