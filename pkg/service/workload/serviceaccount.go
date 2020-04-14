package workload

// ServiceAccount the kubernetes native service account
type ServiceAccount struct {
	WorkloadsResourceHandler
}

// NewServiceAccount exported
func NewServiceAccount() *ServiceAccount {
	return &ServiceAccount{&defaultImplWorkloadsResourceHandler{}}
}
