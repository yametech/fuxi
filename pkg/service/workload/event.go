package workload

type Event struct {
	WorkloadsResourceHandler
}

// NewEvent exported
func NewEvent() *Event {
	return &Event{&defaultImplWorkloadsResourceHandler{}}
}
