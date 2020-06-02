package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ReplicaSet is kubernetes default resource replicaset
type ReplicaSet struct {
	common.WorkloadsResourceHandler // extended for workloadsResourceHandler
}

// NewReplicaSet exported
func NewReplicaSet() *ReplicaSet {
	return &ReplicaSet{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceReplicaSet,
	}}
}
