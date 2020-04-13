package workload

// Ingress the kubernetes native resource ingress
type Ingress struct {
	WorkloadsResourceHandler
}

// NewIngress exported
func NewIngress() *Ingress {
	return &Ingress{&defaultImplWorkloadsResourceHandler{}}
}
