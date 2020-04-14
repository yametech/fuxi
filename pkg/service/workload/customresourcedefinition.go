package workload

// CustomResourceDefinition the kubernetes native CustomResourceDefinition
type CustomResourceDefinition struct {
	WorkloadsResourceHandler
}

// NewCustomResourceDefinition exported
func NewCustomResourceDefinition() *CustomResourceDefinition {
	return &CustomResourceDefinition{&defaultImplWorkloadsResourceHandler{}}
}
