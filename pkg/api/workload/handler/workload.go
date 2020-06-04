package handler

import (
	"encoding/json"
	"fmt"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	nuwav1 "github.com/yametech/nuwa/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
	"strings"

	"github.com/yametech/fuxi/pkg/api/common"
	constraint "github.com/yametech/fuxi/util/common"

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

type deployMetadataItems []deployMetadataItem

type deployMetadataItem struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	Resource        struct {
		Limits struct {
			CPU    int `json:"cpu"`
			Memory int `json:"memory"`
		} `json:"limits"`
		Requests struct {
			CPU    int `json:"cpu"`
			Memory int `json:"memory"`
		} `json:"requests"`
	} `json:"resource"`
	VolumeMounts []interface{} `json:"volumeMounts"`
	Command      []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"command"`
	Args []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"args"`
	OneEnvConfig []struct {
		ID             string `json:"id"`
		Type           string `json:"type"`
		ConfigName     string `json:"configName"`
		ConfigKey      string `json:"configKey"`
		ConfigType     string `json:"configType"`
		EnvConfigName  string `json:"envConfigName"`
		EnvConfigValue string `json:"envConfigValue"`
	} `json:"oneEnvConfig"`
	MultipleEnvConfig []interface{} `json:"multipleEnvConfig"`
	ReadyProbe        struct {
		Status     bool `json:"status"`
		Timeout    int  `json:"timeout"`
		Cycle      int  `json:"cycle"`
		RetryCount int  `json:"retryCount"`
		Delay      int  `json:"delay"`
		Pattern    struct {
			Type     string `json:"type"`
			HTTPPort int    `json:"httpPort"`
			URL      string `json:"url"`
			TCPPort  string `json:"tcpPort"`
			Command  string `json:"command"`
		} `json:"pattern"`
	} `json:"readyProbe"`
	AliveProbe struct {
		Status     bool `json:"status"`
		Timeout    int  `json:"timeout"`
		Cycle      int  `json:"cycle"`
		RetryCount int  `json:"retryCount"`
		Delay      int  `json:"delay"`
		Pattern    struct {
			Type     string `json:"type"`
			HTTPPort int    `json:"httpPort"`
			URL      string `json:"url"`
			TCPPort  string `json:"tcpPort"`
			Command  string `json:"command"`
		} `json:"pattern"`
	} `json:"aliveProbe"`
	LifeCycle struct {
		Status    bool `json:"status"`
		PostStart struct {
			Type     string `json:"type"`
			HTTPPort int    `json:"httpPort"`
			URL      string `json:"url"`
			TCPPort  int    `json:"tcpPort"`
			Command  string `json:"command"`
		} `json:"postStart"`
		PreStop struct {
			Type     string `json:"type"`
			HTTPPort int    `json:"httpPort"`
			URL      string `json:"url"`
			TCPPort  int    `json:"tcpPort"`
			Command  string `json:"command"`
		} `json:"preStop"`
	} `json:"lifeCycle"`
}

type deployTemplate struct {
	AppName   string `json:"appName"`
	Namespace struct {
		Value string `json:"value"`
	} `json:"namespace"`
	TemplateName string `json:"templateName"`
}

func toPodContainer(deployItems ...deployMetadataItem) []corev1.Container {
	contaiers := make([]corev1.Container, len(deployItems), len(deployItems))
	for _, item := range deployItems {
		container := corev1.Container{
			Name:            item.Name,
			Image:           item.Image,
			ImagePullPolicy: corev1.PullPolicy(item.ImagePullPolicy),
			Resources: corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					corev1.ResourceName("cpu"):    resource.MustParse(fmt.Sprintf("%d", item.Resource.Limits.CPU)),
					corev1.ResourceName("memory"): resource.MustParse(fmt.Sprintf("%d", item.Resource.Limits.Memory)),
				},
				Requests: corev1.ResourceList{
					corev1.ResourceName("cpu"):    resource.MustParse(fmt.Sprintf("%d", item.Resource.Requests.CPU)),
					corev1.ResourceName("memory"): resource.MustParse(fmt.Sprintf("%d", item.Resource.Requests.Memory)),
				},
			},
		}
		// Command
		for _, cmd := range item.Command {
			container.Command = append(container.Command, cmd.Value)
		}
		// Args
		for _, arg := range item.Args {
			container.Args = append(container.Command, arg.Value)
		}
		// OneEnvConfig TODO
		for _, evnConfig := range item.OneEnvConfig {
			_ = evnConfig
		}
		contaiers = append(contaiers, container)
	}
	return contaiers
}

func (w *WorkloadsAPI) Deploy(g *gin.Context) {
	deployTemplate := &deployTemplate{}
	if err := g.BindJSON(deployTemplate); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	obj, err := w.workloadsTemplate.Get(constraint.WorkloadsDeployTemplateNamespace, deployTemplate.AppName)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	workloadsTemplate := obj.(*v1.Workloads)

	deployItems := make(deployMetadataItems, 0)
	if err := json.Unmarshal([]byte(*workloadsTemplate.Spec.Metadata), deployItems); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	var runtimeObj runtime.Object
	objectMetadata := metav1.ObjectMeta{
		Name:      deployTemplate.AppName,
		Namespace: deployTemplate.Namespace.Value,
	}
	switch *workloadsTemplate.Spec.ResourceType {
	case "Stone":
		coordinates := make([]nuwav1.CoordinatesGroup, 0)
		runtimeObj = &nuwav1.Stone{
			ObjectMeta: objectMetadata,
			Spec: nuwav1.StoneSpec{
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Name:   deployTemplate.AppName,
						Labels: map[string]string{"app": "deployTemplate.AppName"},
					},
					Spec: corev1.PodSpec{
						Containers: toPodContainer(deployItems...),
					},
				},
				Strategy:    "Release", // TODO
				Coordinates: coordinates,
				Service: corev1.ServiceSpec{
					Type: corev1.ServiceType("ClusterIP"),
				},
			},
		}
	case "Water":
	case "Deployment":
	case "Statefulset":
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(runtimeObj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	unstructuredData := &unstructured.Unstructured{Object: unstructuredObj}

	newObj, err := w.generic.Apply(deployTemplate.Namespace.Value, deployTemplate.AppName, unstructuredData)
	if err != nil {
		common.ToInternalServerError(g, unstructuredData, err)
		return
	}
	g.JSON(
		http.StatusOK,
		[]unstructured.Unstructured{
			*newObj,
		})

	return
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
