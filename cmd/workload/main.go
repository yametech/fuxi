package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/yametech/fuxi/pkg/api/workload/handler"
	"github.com/yametech/fuxi/pkg/preinstall"
	"github.com/yametech/fuxi/pkg/service/common"

	// swagger doc
	file "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"
	_ "github.com/yametech/fuxi/cmd/workload/docs"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger base
// @contact.name laik author
// @contact.url  github.com/yametech
// @contact.email laik.lj@me.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

const (
	name = "go.micro.api.workload"
	ver  = "latest"
)

func initNeed() (web.Service, *gin.Engine, *gin.RouterGroup, *handler.WorkloadsAPI) {
	service, apiInstallConfigure, err := preinstall.InitApi(50, name, ver, "")
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	common.SharedK8sClient = &apiInstallConfigure.DefaultInstallConfigure
	handler.CreateSharedSessionManager(apiInstallConfigure.ClientV1, apiInstallConfigure.RestConfig)

	return service, router, router.Group("/workload"), handler.NewWorkladAPI()
}

var service, router, group, workloadsAPI = initNeed()

func WrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers,X-Access-Token,XKey,Authorization")

		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// #api
	// #v1
	// Pod
	{
		serveHttp := WrapH(handler.CreateAttachHandler("/workload/shell/pod"))
		router.GET("/workload/shell/pod/*path", serveHttp)
		group.GET("/attach/namespace/:namespace/pod/:name/container/:container/:shelltype", PodAttach)
		group.GET("/api/v1/pods", PodList)
		group.GET("/api/v1/namespaces/:namespace/pods", PodList)
		group.GET("/api/v1/namespaces/:namespace/pods/:name", PodGet)
		group.GET("/api/v1/namespaces/:namespace/pods/:name/log", PodLog)
	}

	// Node
	{
		group.GET("/api/v1/nodes", NodeList)
		group.GET("/api/v1/nodes/:node", NodeGet)
		group.POST("/node/annotation/geo", NodeGeoAnnotate)
	}

	// PersistentVolume
	{
		group.GET("/api/v1/persistentvolumes", PersistentVolumeList)
		group.GET("/api/v1/persistentvolumes/:name", PersistentVolumeGet)
		group.DELETE("/api/v1/persistentvolumes/:name", PersistentVolumeDelete)
	}

	// PersistentVolumeClaims
	{
		group.GET("/api/v1/persistentvolumeclaims", PersistentVolumeClaimsList)
		group.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims", PersistentVolumeClaimsList)
		group.GET("/api/v1/namespaces/:namespace/persistentvolumeclaims/:name", PersistentVolumeClaimsGet)
	}

	// Event
	{
		group.GET("/api/v1/events", EventList)
		group.GET("/api/v1/namespaces/:namespace/events", EventList)
		group.GET("/api/v1/namespaces/:namespace/events/:name", EventGet)
	}

	// ResourceQuota
	{
		group.GET("/api/v1/resourcequotas", ResourceQuotaList)
		group.GET("/api/v1/namespaces/:namespace/resourcequotas", ResourceQuotaList)
		group.GET("/api/v1/namespaces/:namespace/resourcequotas/:name", ResourceQuotaGet)
		group.POST("/api/v1/namespaces/:namespace/resourcequotas", workloadsAPI.Apply)
	}

	// Service
	{
		group.GET("/api/v1/services", ServiceList)
		group.GET("/api/v1/namespaces/:namespace/services", ServiceList)
		group.GET("/api/v1/namespaces/:namespace/services/:name", ServiceGet)
	}

	// Endpoint
	{
		group.GET("/api/v1/endpoints", EndpointList)
		group.GET("/api/v1/namespaces/:namespace/endpoints", EndpointList)
		group.GET("/api/v1/namespaces/:namespace/endpoints/:name", EndpointGet)
	}

	// ServiceAccount
	{
		group.GET("/api/v1/serviceaccounts", ServiceAccountList)
		group.GET("/api/v1/namespaces/:namespace/serviceaccounts/:name", ServiceAccountGet)
		group.POST("/serviceaccount/patch/:method", ServiceAccountPatchSecret)
		//group.POST("/api/v1/namespaces/:namespace/serviceaccounts", workloadsAPI.Apply)
	}

	// ConfigMaps
	{
		group.GET("/api/v1/configmaps", ConfigMapsList)
		group.GET("/api/v1/namespaces/:namespace/configmaps", ConfigMapsList)
		group.GET("/api/v1/namespaces/:namespace/configmaps/:name", ConfigMapsGet)
		group.POST("/api/v1/namespaces/:namespace/configmaps", ConfigMapsCreate)
	}

	// Secret
	{
		group.GET("/api/v1/secrets", SecretList)
		group.GET("/api/v1/namespaces/:namespace/secrets", SecretList)
		group.GET("api/v1/ops-secrets", OpsSecretList)
		group.GET("api/v1/namespaces/:namespace/ops-secrets", OpsSecretList)
		group.GET("/api/v1/namespaces/:namespace/secrets/:name", SecretGet)

		group.POST("/api/v1/namespaces/:namespace/secrets", SecretCreate)
		group.POST("/api/v1/namespaces/:namespace/ops-secrets", SecretCreate)

		group.PUT("/api/v1/namespaces/:namespace/secrets/:name", SecretUpdate)
		group.PUT("/api/v1/namespaces/:namespace/ops-secrets/:name", SecretUpdate)
		group.PUT("/api/v1/namespaces/:namespace/secrets", SecretUpdate)
		group.PUT("/api/v1/namespaces/:namespace/ops-secrets", SecretUpdate)
	}

	// #apis
	// #apps v1
	// Deployment
	{
		group.GET("/apis/apps/v1/deployments", DeploymentList)
		group.GET("/apis/apps/v1/namespaces/:namespace/deployments", DeploymentList)
		group.GET("/apis/apps/v1/namespaces/:namespace/deployments/:name", DeploymentGet)
		// deployment scale
		group.GET("/apis/apps/v1/namespaces/:namespace/deployments/:name/scale", workloadsAPI.GetDeploymentScale)
		group.PUT("/apis/apps/v1/namespaces/:namespace/deployments/:name/scale", workloadsAPI.PutDeploymentScale)
	}

	// ReplicaSet
	{
		group.GET("/apis/apps/v1/replicasets", ReplicaSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/replicasets", ReplicaSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/replicasets/:name", ReplicaSetGet)
	}

	// StatefulSet
	{
		group.GET("/apis/apps/v1/statefulsets", StatefulSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/statefulsets", StatefulSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/statefulsets/:name", StatefulSetGet)
	}

	// Stone
	{
		group.GET("/apis/nuwa.nip.io/v1/stones", StoneList)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/stones", StoneList)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/stones/:name", StoneGet)
	}

	// Water
	{
		group.GET("/apis/nuwa.nip.io/v1/waters", WaterList)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/waters", WaterList)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/waters/:name", WaterGet)
	}

	{
		group.GET("/apis/nuwa.nip.io/v1/injectors", InjectorList)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/injectors", InjectorList)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/injectors/:name", InjectorGet)
	}

	// StatefulSet1
	{
		group.GET("/apis/nuwa.nip.io/v1/statefulsets", StatefulSet1List)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/statefulsets", StatefulSet1List)
		group.GET("/apis/nuwa.nip.io/v1/namespaces/:namespace/statefulsets/:name", StatefulSet1Get)
	}

	// DaemonSet
	{
		group.GET("/apis/apps/v1/daemonsets", DaemonSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/daemonsets", DaemonSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/daemonsets/:name", DaemonSetGet)
	}

	// #batch
	// #v1beta1

	// CronJob
	{
		group.GET("/apis/batch/v1beta1/cronjobs", CronJobList)
		group.GET("/apis/batch/v1beta1/namespaces/:namespace/cronjobs", CronJobList)
		group.GET("/apis/batch/v1beta1/namespaces/:namespace/cronjobs/:name", CronJobGet)
	}

	// #v1

	// Job
	{
		group.GET("/apis/batch/v1/jobs", JobList)
		group.GET("/apis/batch/v1/namespaces/:namespace/jobs", JobList)
		group.GET("/apis/batch/v1/namespaces/:namespace/jobs/:name", JobGet)
	}

	// #extensions
	// #v1beta1

	// Ingress
	{
		group.GET("/apis/extensions/v1beta1/ingresses", IngressList)
		group.GET("/apis/extensions/v1beta1/namespaces/:namespace/ingresses", IngressList)
		group.GET("/apis/extensions/v1beta1/namespaces/:namespace/ingresses/:name", IngressGet)
		group.POST("/apis/extensions/v1beta1/namespaces/:namespace/ingresses", workloadsAPI.Apply)
	}

	// #networking.k8s.io
	// #v1

	// NetworkPolicy
	{
		group.GET("/apis/networking.k8s.io/v1/networkpolicies", NetworkPolicyList)
		group.GET("/apis/networking.k8s.io/v1/namespaces/:namespace/networkpolicies", NetworkPolicyList)
		group.GET("/apis/networking.k8s.io/v1/namespaces/:namespace/networkpolicies/:name", NetworkPolicyGet)
	}

	// #storage.k8s.io
	// #v1
	{
		group.GET("/apis/storage.k8s.io/v1/storageclasses", StorageClassList)
		group.GET("/apis/storage.k8s.io/v1/storageclasses/:name", StorageClassGet)
		group.POST("/apis/storage.k8s.io/v1/storageclasses", workloadsAPI.Apply)
	}

	// #autoscaling
	// #v2beta1
	// HorizontalPodAutoscaler
	{
		group.GET("/apis/autoscaling/v2beta1/horizontalpodautoscalers", HorizontalPodAutoscalerList)
		group.GET("/apis/autoscaling/v2beta1/namespaces/:namespace/horizontalpodautoscalers", HorizontalPodAutoscalerList)
		group.GET("/apis/autoscaling/v2beta1/namespaces/:namespace/horizontalpodautoscalers/:name", HorizontalPodAutoscalerGet)
	}

	// kubernetes RBAC
	// #rbac.authorization.k8s.io
	// v1
	// Roles
	{
		group.GET("/apis/rbac.authorization.k8s.io/v1/roles", RoleList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/roles", RoleList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/roles/:name", RoleGet)
	}
	// #rbac.authorization.k8s.io
	// #v1
	// Clusterroles
	{
		group.GET("/apis/rbac.authorization.k8s.io/v1/clusterroles", ClusterRoleList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/clusterroles/:name", ClusterRoleGet)
		group.POST("/apis/rbac.authorization.k8s.io/v1/clusterroles", workloadsAPI.Apply)
	}

	// RoleBinding
	{
		group.GET("/apis/rbac.authorization.k8s.io/v1/rolebindings", RoleBindingList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/rolebindings/:name", RoleBindingGet)
		group.POST("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/rolebindings", workloadsAPI.Apply)
	}

	// ClusterRolebind
	{
		group.GET("/apis/rbac.authorization.k8s.io/v1/clusterrolebindings", ClusterRoleBindList)
		group.GET("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/clusterrolebindings/:name ", ClusterRoleBindGet)
		group.POST("/apis/rbac.authorization.k8s.io/v1/namespaces/:namespace/clusterrolebindings", workloadsAPI.Apply)
	}

	// #policy
	// #v1beta1
	// #podsecuritypolicies
	{
		group.GET("/apis/policy/v1beta1/podsecuritypolicies", PodSecurityPolicieList)
		group.GET("/apis/policy/v1beta1/namespaces/:namespace/podsecuritypolicies/:name", PodSecurityPolicieGet)
		group.POST("/apis/policy/v1beta1/namespaces/:namespace/podsecuritypolicies", workloadsAPI.Apply)
	}

	// #ovn
	// #v1
	// #ips
	{
		group.GET("/apis/kubeovn.io/v1/ips", IPList)
		group.GET("/apis/kubeovn.io/v1/namespaces/:namespace/ips/:name", IPGet)
	}

	// #subnets
	{
		group.GET("/apis/kubeovn.io/v1/subnets", SubNetList)
		group.GET("/apis/kubeovn.io/v1/namespaces/:namespace/subnets/:name", SubNetGet)
		group.POST("apis/kubeovn.io/v1/subnets", workloadsAPI.Apply)
	}

	// #apiextensions.k8s.io/v1beta1
	// #v1beta1
	// CustomResourceDefinition
	{
		group.GET("/apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions", CustomResourceDefinitionList)

		ignores := []string{
			"fuxi.nip.io/v1/workloads",
			"fuxi.nip.io/v1/fields",
			"fuxi.nip.io/v1/forms",
			"fuxi.nip.io/v1/pages",
			"fuxi.nip.io/v1/basedepartments",
			"fuxi.nip.io/v1/baseroles",
			"fuxi.nip.io/v1/baseusers",
			"fuxi.nip.io/v1/tektongraphs",
			"fuxi.nip.io/v1/tektonwebhooks",
			"fuxi.nip.io/v1/tektonstores",
			"nuwa.nip.io/v1/statefulsets",
			"nuwa.nip.io/v1/stones",
			"nuwa.nip.io/v1/waters",
			"nuwa.nip.io/v1/injectors",
			"tekton.dev/v1alpha1/pipelines",
			"tekton.dev/v1alpha1/pipelineruns",
			"tekton.dev/v1alpha1/pipelineresources",
			"tekton.dev/v1alpha1/tasks",
			"tekton.dev/v1alpha1/taskruns",
			"kubeovn.io/v1/ips",
			"kubeovn.io/v1/subnets",
		}
		apiVersions, err := workloadsAPI.ListCustomResourceRouter(ignores)
		if err != nil {
			panic(err)
		}
		routerPath := "/apis/%s"
		for _, apiVersion := range apiVersions {
			relativePath := fmt.Sprintf(routerPath, apiVersion)
			group.GET(relativePath, GeneralCustomResourceDefinitionList)
		}
	}

	// Namespace
	{
		group.GET("/api/v1/namespaces", NamespaceList)
		group.GET("/api/v1/namespaces/:namespace", NamespaceGet)
		group.POST("/api/v1/namespaces", NamespaceCreate)
		group.DELETE("/api/v1/namespaces/:namespace", NamespaceDelete)
		group.POST("/namespaces/annotation/node", NamespacePatchAnnotateNode)
		group.POST("/namespaces/annotation/storageclass", NamespacePatchAnnotateStorageClass)
	}

	// post  workload/stack
	{
		// all resource apply
		group.POST("/stack", workloadsAPI.Apply)

		// workloads deploy post request
		group.POST("/deploy", workloadsAPI.Deploy)
		// other resource  api/apis resource
		group.DELETE("/api/v1/namespaces/:namespace/:resource/:name", workloadsAPI.Delete)
		group.DELETE("/apis/:group/:version/:namespaces_or_resource/:namespace", workloadsAPI.Delete)
		group.DELETE("/apis/:group/:version/:namespaces_or_resource/:namespace/:resource/:name", workloadsAPI.Delete)
	}

	// fuxi.nip.io
	// Workloads
	{
		group.GET("/apis/fuxi.nip.io/v1/workloads", WorkloadsTemplateList)
		// /apis/fuxi.nip.io/v1/namespaces/dxp/workloads
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/workloads", WorkloadsTemplateListSharedNamespace)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/workloads/:name", WorkloadsTemplateGet)
		group.POST("/apis/fuxi.nip.io/v1/workloads", WorkloadsTemplateCreate)
	}

	// Field
	{
		group.GET("/apis/fuxi.nip.io/v1/fields", FieldList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/fields/:name", FieldGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/fields", FieldCreate)
	}

	// Form
	{
		group.GET("/apis/fuxi.nip.io/v1/forms", FormList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/forms/:name", FormGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/forms", FormCreate)
	}

	// Page
	{
		group.GET("/apis/fuxi.nip.io/v1/pages", PageList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/pages/:name", PageGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/pages", PageCreate)
	}

	// Metrics
	{
		group.POST("/metrics", workloadsAPI.Metrics)
		group.GET("/apis/metrics.k8s.io/v1beta1/nodes", workloadsAPI.NodeMetrics)

		group.GET("/apis/metrics.k8s.io/v1beta1/pods", workloadsAPI.PodMetricsList)
		// GET /workload/apis/metrics.k8s.io/v1beta1/namespaces/rook-ceph/pods
		group.GET("/apis/metrics.k8s.io/v1beta1/namespaces/:namespace/pods", workloadsAPI.PodMetrics)
	}

	{
		group.GET("/config", func(g *gin.Context) {
			g.JSON(http.StatusOK, gin.H{
				"lensVersion":       "1.0",
				"lensTheme":         "",
				"userName":          "admin",
				"token":             "",
				"allowedNamespaces": "[]",
				"isClusterAdmin":    true,
				"chartEnable":       true,
				"kubectlAccess":     true,
			})
		})
	}

	// tekton.dev
	// v1alpha1
	{
		// pipeline
		group.GET("/apis/tekton.dev/v1alpha1/pipelines", PipelineList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelines", PipelineList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelines/:name", PipelineGet)
		group.POST("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelines", workloadsAPI.Apply)
		group.PUT("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelines/:name", workloadsAPI.Apply)

		// pipelineRun
		group.GET("/apis/tekton.dev/v1alpha1/pipelineruns", PipelineRunList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineruns", PipelineRunList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineruns/:name", PipelineRunGet)
		group.POST("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineruns", workloadsAPI.Apply)

		// task
		group.GET("/apis/tekton.dev/v1alpha1/tasks", TaskList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/tasks", TaskList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/tasks/:name", TaskGet)
		group.POST("/apis/tekton.dev/v1alpha1/namespaces/:namespace/tasks", workloadsAPI.Apply)

		// taskRun
		group.GET("/apis/tekton.dev/v1alpha1/taskruns", TaskRunList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/taskruns", TaskRunList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/taskruns/:name", TaskRunGet)

		// pipelineResource
		group.GET("/apis/tekton.dev/v1alpha1/pipelineresources", PipelineResourceList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineresources", PipelineResourceList)
		group.GET("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineresources/:name", PipelineResourceGet)
		group.POST("/apis/tekton.dev/v1alpha1/namespaces/:namespace/pipelineresources", workloadsAPI.Apply)

	}

	// TektonGraph
	{
		group.GET("/apis/fuxi.nip.io/v1/tektongraphs", TektonGraphList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektongraphs", TektonGraphList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektongraphs/:name", TektonGraphGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektongraphs", TektonGraphCreate)
	}

	// TektonStore
	{
		group.GET("/apis/fuxi.nip.io/v1/tektonstores", TektonStoreList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektonstores", TektonStoreList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektonstores/:name", TektonStoreGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektonstores", TektonStoreCreate)
	}

	// TektonWebHook
	{
		group.GET("/apis/fuxi.nip.io/v1/tektonwebhooks", TektonWebHookList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektonwebhooks", TektonWebHookList)
		group.GET("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektonwebhooks/:name", TektonWebHookGet)
		group.POST("/apis/fuxi.nip.io/v1/namespaces/:namespace/tektonwebhooks", TektonWebHookCreate)
	}

	// watch the group resource
	{
		group.GET("/watch", WatchStream)
	}

	// Swag
	{
		/// Then, if you set envioment variable DEV_OPEN_SWAGGER to anything, /swagger/*any will respond 404, just like when route unspecified.
		/// Release production environment can be turned on
		group.GET("/swagger/*any", swag.DisablingWrapHandler(file.Handler, "DEV_OPEN_SWAGGER"))
	}

	service.Handle("/", router)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
