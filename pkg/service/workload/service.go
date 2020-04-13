package workload

// Service the kubernetes native resource services
type Service struct {
	WorkloadsResourceHandler
}

// NewService exported
func NewService() *Service {
	return &Service{&defaultImplWorkloadsResourceHandler{}}
}
