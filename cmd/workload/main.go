package main

import (
	//"log"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/yametech/fuxi/pkg/api/workload/handler"
	"github.com/yametech/fuxi/pkg/k8s/client"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/preinstall"
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
	"github.com/yametech/fuxi/thirdparty/lib/wrapper/tracer/opentracing/gin2micro"
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
	service, _, err := preinstall.InitApi(50, name, ver, "")
	if err != nil {
		panic(err)
	}

	hystrix.DefaultTimeout = 5000
	wrapper := hystrixplugin.NewClientWrapper()
	_ = wrapper

	router := gin.Default()
	router.Use(gin2micro.TracerWrapper)

	err = workloadservice.NewK8sClientSet(dyn.SharedCacheInformerFactory, client.K8sResourceHandler, client.RestConf)
	if err != nil {
		panic(err)
	}

	handler.CreateSharedSessionManager()

	return service, router, router.Group("/workload"), handler.NewWorkladAPI()
}

var service, router, group, workloadsAPI = initNeed()

func main() {

	// Pod
	{
		group.GET("/shell/pod", gin.WrapH(handler.CreateAttachHandler("/workload/shell/pod")))
		group.GET("/attach/namespace/:namespace/pod/:name/container/:container", PodAttach)
		group.GET("/api/v1/pods", PodList)
		group.GET("/api/v1/namespace/:namespace/pod/:name", PodGet)
	}

	// Event
	{
		group.GET("/api/v1/events", EventList)
		group.GET("/api/v1/namespace/:namespace/event/:name", EventGet)
	}

	// Node
	{
		group.GET("/api/v1/nodes", NodeList)
		group.GET("/api/v1/nodes/:node", NodeGet)
	}

	// Deployment
	{
		group.GET("/apis/apps/v1/deployments", DeploymentList)
		group.GET("/apis/apps/v1/namespaces/:namespace/deployments/:name", DeploymentGet)
	}

	// CronJob
	{
		group.GET("/apis/batch/v1beta1/cronjobs", CronJobList)
		group.GET("/apis/batch/v1beta1/namespaces/:namespace/cronjobs/:name", CronJobGet)
	}

	// StatefulSet
	{
		group.GET("/apis/apps/v1/statefulsets", StatefulSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/statefulsets/:name", StatefulSetGet)
	}

	// DaemonSet
	{
		group.GET("/apis/apps/v1/daemonsets", DaemonSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/daemonsets/:name", DaemonSetGet)
	}

	// Job
	{
		group.GET("/apis/batch/v1/jobs", JobList) //TODO fix route
		group.GET("/apis/batch/v1/namespaces/:namespace/jobs/:name", JobGet)
	}

	// ReplicaSet
	{
		group.GET("/apis/apps/v1/replicasets", ReplicaSetList)
		group.GET("/apis/apps/v1/namespaces/:namespace/replicasets/:name", ReplicaSetGet)
	}

	// ConfigMaps
	{
		group.GET("/api/v1/configmaps", ConfigMapsList)
		group.GET("/api/v1/namespaces/:namespace/configmaps/:name", ConfigMapsGet)
	}

	// Secret
	{
		group.GET("/api/v1/secrets", SecretList)
		group.GET("/api/v1/namespaces/:namespace/secrets/:name", SecretGet)
	}

	// ResourceQuota
	{
		group.GET("/api/v1/resourcequotas", ResourceQuotaList)
		group.GET("/api/v1/namespaces/:namespace/resourcequotas/:name", ResourceQuotaGet)
	}

	// Service
	{
		group.GET("/api/v1/services", ServiceList)
		group.GET("/api/v1/namespaces/:namespace/services/:name", ServiceGet)
	}

	// Ingress
	{
		group.GET("/apis/extensions/v1beta1/ingresses", IngressList)
		group.GET("/apis/extensions/v1beta1/namespaces/:namespace/ingresses/:name", IngressGet)
	}

	// NetworkPolicy
	{
		group.GET("/apis/networking.k8s.io/v1/networkpolicies", NetworkPolicyList)
		group.GET("/apis/networking.k8s.io/v1/namespaces/:namespace/networkpolicies/:name", NetworkPolicyGet)
	}

	// HorizontalPodAutoscaler
	{
		group.GET("/apis/autoscaling/v2beta1/horizontalpodautoscalers", HorizontalPodAutoscalerList)
		group.GET("/apis/autoscaling/v2beta1/namespaces/:namespace/horizontalpodautoscalers/:name", HorizontalPodAutoscalerGet)
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
