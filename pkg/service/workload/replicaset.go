package workload

import dyn "github.com/yametech/fuxi/pkg/kubernetes/client"

// ReplicaSet is kubernetes default resource replicaset
type ReplicaSet struct {
	WorkloadsResourceHandler // extended for workloadsResourceHandler
}

// NewReplicaSet exported
func NewReplicaSet() *ReplicaSet {
	return &ReplicaSet{&defaultImplWorkloadsResourceHandler{
		dyn.ResourceReplicaSet,
	}}
}
