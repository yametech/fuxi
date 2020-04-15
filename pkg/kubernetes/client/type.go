package client

import "k8s.io/apimachinery/pkg/runtime/schema"

type ResourceName string

// nuwa resource
const (
	Water                    ResourceName = "water"
	Deployment               ResourceName = "deployment"
	Stone                    ResourceName = "stone"
	StatefulSet              ResourceName = "statefulset"
	DaemonSet                ResourceName = "daemonsets"
	Injector                 ResourceName = "injector"
	Pod                      ResourceName = "pod"
	Job                      ResourceName = "jobs"
	CronJobs                 ResourceName = "cronjobs"
	ReplicaSet               ResourceName = "replicasets"
	Event                    ResourceName = "events"
	Node                     ResourceName = "nodes"
	ConfigMaps               ResourceName = "configmaps"
	Secrets                  ResourceName = "secrets"
	ResourceQuota            ResourceName = "resourcequotas"
	Service                  ResourceName = "services"
	Ingress                  ResourceName = "ingresses"
	NetworkPolicy            ResourceName = "networkpolicies"
	HorizontalPodAutoscaler  ResourceName = "horizontalpodautoscalers"
	CustomResourceDefinition ResourceName = "customresourcedefinitions"
	PersistentVolume         ResourceName = "persistentvolumes"
	PersistentVolumeClaims   ResourceName = "persistentvolumeclaims"
	StorageClass             ResourceName = "storageclasses"
	ServiceAccount           ResourceName = "serviceaccounts"
	Role                     ResourceName = "roles"
	RoleBinding              ResourceName = "rolebindings"
)

// GroupVersionResources describe resource collection
var GroupVersionResources = map[ResourceName]schema.GroupVersionResource{
	Water:                    {Group: "nuwa.nip.io", Version: "v1", Resource: "waters"},
	Deployment:               {Group: "apps", Version: "v1", Resource: "deployments"},
	Stone:                    {Group: "nuwa.nip.io", Version: "v1", Resource: "stones"},
	StatefulSet:              {Group: "apps", Version: "v1", Resource: "statefulsets"},
	DaemonSet:                {Group: "apps", Version: "v1", Resource: "daemonsets"},
	Injector:                 {Group: "nuwa.nip.io", Version: "v1", Resource: "injectors"},
	Pod:                      {Group: "", Version: "v1", Resource: "pods"},
	Node:                     {Group: "", Version: "v1", Resource: "nodes"},
	Event:                    {Group: "", Version: "v1", Resource: "events"},
	Job:                      {Group: "batch", Version: "v1", Resource: "jobs"},
	CronJobs:                 {Group: "batch", Version: "v1beta1", Resource: "cronjobs"},
	ReplicaSet:               {Group: "apps", Version: "v1", Resource: "replicasets"},
	ConfigMaps:               {Group: "", Version: "v1", Resource: "configmaps"},
	Secrets:                  {Group: "", Version: "v1", Resource: "secrets"},
	ResourceQuota:            {Group: "", Version: "v1", Resource: "resourcequotas"},
	Service:                  {Group: "", Version: "v1", Resource: "services"},
	Ingress:                  {Group: "extensions", Version: "v1beta1", Resource: "ingresses"},
	NetworkPolicy:            {Group: "networking.k8s.io", Version: "v1", Resource: "networkpolicies"},
	HorizontalPodAutoscaler:  {Group: "autoscaling", Version: "v2beta1", Resource: "horizontalpodautoscalers"},
	CustomResourceDefinition: {Group: "apiextensions.k8s.io", Version: "v1beta1", Resource: "customresourcedefinitions"},
	PersistentVolume:         {Group: "", Version: "v1", Resource: "persistentvolumes"},
	PersistentVolumeClaims:   {Group: "", Version: "v1", Resource: "persistentvolumeclaims"},
	StorageClass:             {Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"},
	ServiceAccount:           {Group: "", Version: "v1", Resource: "serviceaccounts"},
	Role:                     {Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "roles"},
	RoleBinding:              {Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "rolebindings"},
}

var (
	ResourceWater                    = GetGVR(Water)
	ResourceJob                      = GetGVR(Job)
	ResourceDeployment               = GetGVR(Deployment)
	ResourceStone                    = GetGVR(Stone)
	ResourceStatefulSet              = GetGVR(StatefulSet)
	ResourceDaemonSet                = GetGVR(DaemonSet)
	ResourceInjector                 = GetGVR(Injector)
	ResourcePod                      = GetGVR(Pod)
	ResourceCronJobs                 = GetGVR(CronJobs)
	ResourceReplicaSet               = GetGVR(ReplicaSet)
	ResourceEvent                    = GetGVR(Event)
	ResourceNode                     = GetGVR(Node)
	ResourceConfigMaps               = GetGVR(ConfigMaps)
	ResourceSecrets                  = GetGVR(Secrets)
	ResourceResourceQuota            = GetGVR(ResourceQuota)
	ResourceService                  = GetGVR(Service)
	ResourceIngress                  = GetGVR(Ingress)
	ResourceNetworkPolicy            = GetGVR(NetworkPolicy)
	ResourceHorizontalPodAutoscaler  = GetGVR(HorizontalPodAutoscaler)
	ResourceCustomResourceDefinition = GetGVR(CustomResourceDefinition)
	ResourcePersistentVolume         = GetGVR(PersistentVolume)
	ResourcePersistentVolumeClaims   = GetGVR(PersistentVolumeClaims)
	ResourceStorageClass             = GetGVR(StorageClass)
	ResourceServiceAccount           = GetGVR(ServiceAccount)
	ResourceRole                     = GetGVR(Role)
	ResourceRoleBinding              = GetGVR(RoleBinding)
)

func GetGVR(rs ResourceName) schema.GroupVersionResource {
	gvr, exist := GroupVersionResources[rs]
	if !exist {
		panic("try to get an undefined resource")
	}
	return gvr
}
