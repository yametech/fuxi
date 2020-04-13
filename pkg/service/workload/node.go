package workload

type Node struct {
	WorkloadsResourceHandler
}

// NewNode exported
func NewNode() *Node {
	return &Node{&defaultImplWorkloadsResourceHandler{}}
}
