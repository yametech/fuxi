package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

type Node struct {
	common.WorkloadsResourceHandler
}

// NewNode exported
func NewNode() *Node {
	return &Node{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceNode,
	}}
}
