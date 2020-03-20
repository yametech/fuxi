package client

import "k8s.io/apimachinery/pkg/runtime/schema"

type ResourceName string

// nuwa resource
const (
	Water       ResourceName = "water"
	Deployment  ResourceName = "deployment"
	Stone       ResourceName = "stone"
	StatefulSet ResourceName = "statefulset"
	Injector    ResourceName = "injector"
	Pod         ResourceName = "pod"
)

var GroupVersionResources = map[ResourceName]schema.GroupVersionResource{
	Water:       {Group: "nuwa.nip.io", Version: "v1", Resource: "waters"},
	Deployment:  {Group: "apps", Version: "v1", Resource: "deployments"},
	Stone:       {Group: "nuwa.nip.io", Version: "v1", Resource: "stones"},
	StatefulSet: {Group: "nuwa.nip.io", Version: "v1", Resource: "statefulsets"},
	Injector:    {Group: "nuwa.nip.io", Version: "v1", Resource: "injectors"},
	Pod:         {Group: "", Version: "v1", Resource: "pods"},
}

var (
	ResourceWater       = gvr(Water)
	ResourceDeployment  = gvr(Deployment)
	ResourceStone       = gvr(Stone)
	ResourceStatefulSet = gvr(StatefulSet)
	ResourceInjector    = gvr(Injector)
	ResourcePod         = gvr(Pod)
)

func gvr(rs ResourceName) schema.GroupVersionResource {
	gvr, exist := GroupVersionResources[rs]
	if !exist {
		panic("try to get an undefined resource")
	}
	return gvr
}
