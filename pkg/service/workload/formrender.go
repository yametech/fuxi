package workload

type FormRender struct {
	WorkloadsResourceHandler
}

// NewFormRender exported
func NewFormRender() *FormRender {
	return &FormRender{&defaultImplWorkloadsResourceHandler{}}
}
