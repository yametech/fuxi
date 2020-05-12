package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

type Node struct {
	WorkloadsResourceHandler
}

// NewNode exported
func NewNode() *Node {
	return &Node{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceNode,
	}}
}
