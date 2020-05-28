package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Node struct {
	common.WorkloadsResourceHandler
}

// NewNode exported
func NewNode() *Node {
	return &Node{&common.DefaultImplWorkloadsResourceHandler{
		dyn.ResourceNode,
	}}
}
