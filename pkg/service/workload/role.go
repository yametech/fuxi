package workload

// Role the kubernetes native role
type Role struct {
	WorkloadsResourceHandler
}

// NewRole exported
func NewRole() *Role {
	return &Role{&defaultImplWorkloadsResourceHandler{}}
}
