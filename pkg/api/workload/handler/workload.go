package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
	"strings"
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
	}
}

// /:group/:version/namespaces/:namespace/:resource/:name
func (w *WorkloadsAPI) Delete(g *gin.Context) {
	group := g.Param("group")
	version := g.Param("version")
	namespace := g.Param("namespace")
	resource := g.Param("resource")
	name := g.Param("name")
	if version == "" || namespace == "" || resource == "" || name == "" {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    "",
				status: "Request bad parameter",
			},
		)
		return
	}

	if group == "api" {
		group = ""
	}

	gvr := schema.GroupVersionResource{Group: group, Version: version, Resource: resource}
	if err := w.generic.Delete(gvr, namespace, name); err != nil {
		g.JSON(
			http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   "",
				msg:    err.Error(),
				status: "delete resource internal server error",
			},
		)
		return
	}
	g.JSON(http.StatusOK, nil)
}

func (w *WorkloadsAPI) Apply(g *gin.Context) {
	var formData map[string]interface{}
	if err := g.BindJSON(&formData); err != nil {
		toRequestParamsError(g, err)
		return
	}
	apiVersion := formData["apiVersion"].(string)
	// split apiVersion version and kind
	apiVersions := strings.Split(apiVersion, "/")
	kind, ok := formData["kind"].(string)
	if !ok {
		toRequestParamsError(g, fmt.Errorf("form data kind not define"))
		return
	}
	kind = fmt.Sprintf("%s%s", strings.ToLower(kind), "s")

	md, ok := formData["metadata"]
	if !ok {
		toRequestParamsError(g, fmt.Errorf("form data metadata not define"))
		return
	}

	metadata, ok := md.(map[string]interface{})
	if !ok {
		toRequestParamsError(g, fmt.Errorf("form data metadata type error"))
		return
	}

	namespace, ok := metadata["namespace"].(string)
	if !ok && kind != "namespaces" {
		toRequestParamsError(g, fmt.Errorf("namespace not define"))
		return
	}

	name, ok := metadata["name"].(string)
	if !ok{
		toRequestParamsError(g, fmt.Errorf("name not define"))
		return
	}

	unstructuredData := &unstructured.Unstructured{Object: formData}

	var runtimeClassGVR schema.GroupVersionResource
	if len(apiVersions) > 1 {
		runtimeClassGVR = schema.GroupVersionResource{Group: apiVersions[0], Version: apiVersions[1], Resource: kind}
	} else if len(apiVersions) == 1 {
		runtimeClassGVR = schema.GroupVersionResource{Group: "", Version: apiVersions[0], Resource: kind}
	} else {
		toInternalServerError(g, formData, nil)
		return
	}

	newObj, err := w.generic.Apply(runtimeClassGVR, namespace, name, unstructuredData)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}

	g.JSON(
		http.StatusOK,
		[]unstructured.Unstructured{
			*newObj,
		})
}
