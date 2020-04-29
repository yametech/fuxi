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

	if group != "apps" {
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
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter",
			},
		)
		return
	}
	unstructuredData := &unstructured.Unstructured{Object: formData}
	md, _ := formData["metadata"]
	metadata := md.(map[string]interface{})
	namespace := metadata["namespace"].(string)
	name := metadata["name"].(string)
	apiVersion := formData["apiVersion"].(string)
	apiVersions := strings.Split(apiVersion, "/")

	kind := formData["kind"].(string)
	kind = fmt.Sprintf("%s%s", strings.ToLower(kind), "s")
	gvr := schema.GroupVersionResource{Group: apiVersions[0], Version: apiVersions[1], Resource: kind}
	err := w.generic.Apply(gvr, namespace, name, unstructuredData)
	if err != nil {
		g.JSON(http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   "",
				msg:    err.Error(),
				status: "apply error",
			},
		)
		return
	}

	g.JSON(http.StatusOK, []unstructured.Unstructured{
		*unstructuredData,
	})
}
