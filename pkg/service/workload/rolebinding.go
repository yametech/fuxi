package workload

// RoleBinding the kubernetes native role binding
type RoleBinding struct {
	WorkloadsResourceHandler
}

// NewRoleBinding exported
func NewRoleBinding() *RoleBinding {
	return &RoleBinding{&defaultImplWorkloadsResourceHandler{}}
}
