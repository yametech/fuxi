package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yametech/fuxi/pkg/api/common"

	"github.com/gin-gonic/gin"
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// WorkloadsAPI all resource operate
type WorkloadsAPI struct {
	deployments              *workloadservice.Deployment
	job                      *workloadservice.Job
	cronJob                  *workloadservice.CronJob
	statefulSet              *workloadservice.StatefulSet
	daemonSet                *workloadservice.DaemonSet
	replicaSet               *workloadservice.ReplicaSet
	pod                      *workloadservice.Pod
	event                    *workloadservice.Event
	node                     *workloadservice.Node
	configMaps               *workloadservice.ConfigMaps
	secret                   *workloadservice.Secrets
	resourceQuota            *workloadservice.ResourceQuota
	service                  *workloadservice.Service
	ingress                  *workloadservice.Ingress
	networkPolicy            *workloadservice.NetworkPolicy
	horizontalPodAutoscaler  *workloadservice.HorizontalPodAutoscaler
	customResourceDefinition *workloadservice.CustomResourceDefinition
	persistentVolume         *workloadservice.PersistentVolume
	persistentVolumeClaims   *workloadservice.PersistentVolumeClaims
	storageClass             *workloadservice.StorageClass
	serviceAccount           *workloadservice.ServiceAccount
	role                     *workloadservice.Role
	roleBinding              *workloadservice.RoleBinding
	namespace                *workloadservice.Namespace
	metrics                  *workloadservice.Metrics
	generic                  *workloadservice.Generic
	field                    *workloadservice.Field
	form                     *workloadservice.Form
	page                     *workloadservice.Page
	statefulSet1             *workloadservice.StatefulSet1
	endpoint                 *workloadservice.Endpoint
	clusterrole              *workloadservice.ClusterRole
	clusterRoleBinding       *workloadservice.ClusterRoleBinding
	workloadsTemplate        *workloadservice.WorkloadsTemplate
	pipeline                 *workloadservice.Pipeline
	pipelineRun              *workloadservice.PipelineRun
	task                     *workloadservice.Task
	taskrun                  *workloadservice.TaskRun
	pipelineResource         *workloadservice.PipelineResource
	podsecuritypolicies      *workloadservice.PodSecurityPolicies
}

func NewWorkladAPI() *WorkloadsAPI {
	return &WorkloadsAPI{
		deployments:              workloadservice.NewDeployment(),
		cronJob:                  workloadservice.NewCronJob(),
		statefulSet:              workloadservice.NewStatefulSet(),
		daemonSet:                workloadservice.NewDaemonSet(),
		job:                      workloadservice.NewJob(),
		replicaSet:               workloadservice.NewReplicaSet(),
		pod:                      workloadservice.NewPod(),
		event:                    workloadservice.NewEvent(),
		node:                     workloadservice.NewNode(),
		configMaps:               workloadservice.NewConfigMaps(),
		secret:                   workloadservice.NewSecrets(),
		resourceQuota:            workloadservice.NewResourceQuota(),
		service:                  workloadservice.NewService(),
		ingress:                  workloadservice.NewIngress(),
		networkPolicy:            workloadservice.NewNetworkPolicy(),
		horizontalPodAutoscaler:  workloadservice.NewHorizontalPodAutoscaler(),
		customResourceDefinition: workloadservice.NewCustomResourceDefinition(),
		persistentVolume:         workloadservice.NewPersistentVolume(),
		persistentVolumeClaims:   workloadservice.NewPersistentVolumeClaims(),
		storageClass:             workloadservice.NewStorageClass(),
		serviceAccount:           workloadservice.NewServiceAccount(),
		role:                     workloadservice.NewRole(),
		roleBinding:              workloadservice.NewRoleBinding(),
		namespace:                workloadservice.NewNamespace(),
		metrics:                  workloadservice.NewMetrics(),
		generic:                  workloadservice.NewGeneric(),
		field:                    workloadservice.NewField(),
		form:                     workloadservice.NewForm(),
		page:                     workloadservice.NewPage(),
		statefulSet1:             workloadservice.NewStatefulSet1(),
		endpoint:                 workloadservice.NewEndpoint(),
		clusterrole:              workloadservice.NewClusterRole(),
		clusterRoleBinding:       workloadservice.NewClusterRoleBinding(),
		workloadsTemplate:        workloadservice.NewWorkloadsTemplate(),
		pipeline:                 workloadservice.NewPipeline(),
		pipelineRun:              workloadservice.NewPipelineRun(),
		podsecuritypolicies:      workloadservice.NewPodSecurityPolicies(),
	}
}

// /api/:version/:resource/:name
// /api/:version/namespaces/:namespace/:resource/:name
// /apis/:group/:version/namespaces/:namespace
func (w *WorkloadsAPI) Delete(g *gin.Context) {
	group := g.Param("group")
	version := g.Param("version")

	namespace := g.Param("namespace")
	resource := g.Param("resource")
	name := g.Param("name")

	if namespace == "" || resource == "" || name == "" {
		common.ToRequestParamsError(g, fmt.Errorf("request params not define"))
		return
	}

	if group == "" {
		version = "v1"
	}

	gvr := schema.GroupVersionResource{Group: group, Version: version, Resource: resource}
	w.generic.SetGroupVersionResource(gvr)
	if err := w.generic.Delete(namespace, name); err != nil {
		common.ToInternalServerError(g, "delete resource internal server error", err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

var ignoreList = []string{
	"namespaces",
	"subnets",
	"clusterroles",
	"clusterrolebindings",
}

var in = func(item string) bool {
	for _, ignoreItem := range ignoreList {
		if ignoreItem == item {
			return true
		}
	}
	return false
}

func (w *WorkloadsAPI) Apply(g *gin.Context) {
	var formData map[string]interface{}
	if err := g.BindJSON(&formData); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	apiVersion := formData["apiVersion"].(string)
	// split apiVersion version and kind
	apiVersions := strings.Split(apiVersion, "/")
	kind, ok := formData["kind"].(string)
	if !ok {
		common.ToRequestParamsError(g, fmt.Errorf("form data kind not define"))
		return
	}
	if strings.HasSuffix(strings.ToLower(kind), "ss") {
		// Compatible with ingress resources
		kind = fmt.Sprintf("%s%s", strings.ToLower(kind), "es")
	} else if strings.HasSuffix(strings.ToLower(kind), "s") {
		kind = strings.ToLower(kind)
	} else {
		kind = fmt.Sprintf("%s%s", strings.ToLower(kind), "s")
	}

	md, ok := formData["metadata"]
	if !ok {
		common.ToRequestParamsError(g, fmt.Errorf("form data metadata not define"))
		return
	}

	metadata, ok := md.(map[string]interface{})
	if !ok {
		common.ToRequestParamsError(g, fmt.Errorf("form data metadata type error"))
		return
	}

	namespace, ok := metadata["namespace"].(string)
	// ignore cluster scope resource
	if !ok && !in(kind) {
		common.ToRequestParamsError(g, fmt.Errorf("namespace not define"))
		return
	}

	name, ok := metadata["name"].(string)
	if !ok {
		common.ToRequestParamsError(g, fmt.Errorf("name not define"))
		return
	}

	unstructuredData := &unstructured.Unstructured{Object: formData}

	var runtimeClassGVR schema.GroupVersionResource
	if len(apiVersions) > 1 {
		runtimeClassGVR = schema.GroupVersionResource{Group: apiVersions[0], Version: apiVersions[1], Resource: kind}
	} else if len(apiVersions) == 1 {
		runtimeClassGVR = schema.GroupVersionResource{Group: "", Version: apiVersions[0], Resource: kind}
	} else {
		common.ToInternalServerError(g, formData, nil)
		return
	}

	w.generic.SetGroupVersionResource(runtimeClassGVR)
	newObj, err := w.generic.Apply(namespace, name, unstructuredData)
	if err != nil {
		common.ToInternalServerError(g, runtimeClassGVR.String(), err)
		return
	}
	g.JSON(
		http.StatusOK,
		[]unstructured.Unstructured{
			*newObj,
		})
}
