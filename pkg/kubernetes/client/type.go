package client

import "k8s.io/apimachinery/pkg/runtime/schema"

type ResourceName string

// nuwa resource
const (
	Water       ResourceName = "water"
	Deployment  ResourceName = "deployment"
	ReplicaSet  ResourceName = "replicaset"
	Stone       ResourceName = "stone"
	StatefulSet ResourceName = "statefulset"
	Injector    ResourceName = "injector"
	Pod         ResourceName = "pod"
)

var GroupVersionResources = map[ResourceName]schema.GroupVersionResource{
	Deployment:  {Group: "apps", Version: "v1", Resource: "deployments"},
	ReplicaSet:  {Group: "apps", Version: "v1", Resource: "replicasets"},
	Stone:       {Group: "nuwa.nip.io", Version: "v1", Resource: "stones"},
	StatefulSet: {Group: "nuwa.nip.io", Version: "v1", Resource: "statefulsets"},
	Water:       {Group: "nuwa.nip.io", Version: "v1", Resource: "waters"},
	Injector:    {Group: "nuwa.nip.io", Version: "v1", Resource: "injectors"},
	Pod:         {Group: "core", Version: "v1", Resource: "pods"},
}

var (
	ResourceDeployment  = gvr(Deployment)
	ResourceReplicaSet  = gvr(ReplicaSet)
	ResourceStone       = gvr(Stone)
	ResourceStatefulSet = gvr(StatefulSet)
	ResourcePod         = gvr(Pod)
	ResourceInjector    = gvr(Injector)
)

func gvr(rs ResourceName) schema.GroupVersionResource {
	gvr, exist := GroupVersionResources[rs]
	if !exist {
		panic("try to get an undefined resource")
	}
	return gvr
}
