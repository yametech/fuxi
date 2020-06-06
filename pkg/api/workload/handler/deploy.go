package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	constraint "github.com/yametech/fuxi/util/common"
	nuwav1 "github.com/yametech/nuwa/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
	"strconv"
)

type deployMetadataItems []deployMetadataItem

type deployMetadataItem struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	Resource        struct {
		Limits struct {
			CPU    float64 `json:"cpu"`
			Memory int     `json:"memory"`
		} `json:"limits"`
		Requests struct {
			CPU    float64 `json:"cpu"`
			Memory int     `json:"memory"`
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
			TCPPort  int    `json:"tcpPort"`
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
			TCPPort  int    `json:"tcpPort"`
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
	Replicas     string `json:"replicas"`
	TemplateName string `json:"templateName"`
}

func toPodContainer(deployItems ...deployMetadataItem) []corev1.Container {
	containers := make([]corev1.Container, 0)
	for _, item := range deployItems {
		container := corev1.Container{
			Name:            item.Name,
			Image:           item.Image,
			ImagePullPolicy: corev1.PullPolicy(item.ImagePullPolicy),
			Resources: corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					corev1.ResourceName("cpu"): resource.MustParse(
						fmt.Sprintf("%f", item.Resource.Limits.CPU),
					),
					corev1.ResourceName("memory"): resource.MustParse(fmt.Sprintf("%d", item.Resource.Limits.Memory*1024*1024)),
				},
				Requests: corev1.ResourceList{
					corev1.ResourceName("cpu"):    resource.MustParse(fmt.Sprintf("%f", item.Resource.Requests.CPU)),
					corev1.ResourceName("memory"): resource.MustParse(fmt.Sprintf("%d", item.Resource.Requests.Memory*1024*1024)),
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
		containers = append(containers, container)
	}
	return containers
}

func groupBy(cds nuwav1.Coordinates, replicas int32) (group int, result []nuwav1.CoordinatesGroup) {
	var temp = make(map[string]nuwav1.Coordinates)
	for _, coordinate := range cds {
		if _, ok := temp[coordinate.Zone]; !ok {
			temp[coordinate.Zone] = make(nuwav1.Coordinates, 0)
		}
		temp[coordinate.Zone] = append(temp[coordinate.Zone], coordinate)
	}
	for k, v := range temp {
		result = append(result, nuwav1.CoordinatesGroup{
			Group:    k,
			Zoneset:  v,
			Replicas: &replicas,
		})
	}
	return len(temp), result
}

func (w *WorkloadsAPI) Deploy(g *gin.Context) {
	deployTemplate := &deployTemplate{}
	if err := g.BindJSON(deployTemplate); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	obj, err := w.workloadsTemplate.Get(constraint.WorkloadsDeployTemplateNamespace, deployTemplate.TemplateName)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	workloadsUnstructured := obj.(*unstructured.Unstructured)

	workloads := &v1.Workloads{}
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(workloadsUnstructured.Object, workloads); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	deployItems := make(deployMetadataItems, 0)
	if err := json.Unmarshal([]byte(*workloads.Spec.Metadata), &deployItems); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	var runtimeObj runtime.Object

	var runtimeClassGVR schema.GroupVersionResource
	switch *workloads.Spec.ResourceType {
	case "Stone":
		obj, err := w.namespace.RemoteGet("", deployTemplate.Namespace.Value)
		if err != nil {
			common.ToRequestParamsError(g, err)
			return
		}
		namespaceUnstructured := obj.(*unstructured.Unstructured)
		namespace := &corev1.Namespace{}
		if err = runtime.DefaultUnstructuredConverter.FromUnstructured(namespaceUnstructured.Object, namespace); err != nil {
			common.ToRequestParamsError(g, err)
			return
		}

		notResourceAllocatedError := fmt.Errorf("node resources are not allocated in this namespace, please contact the administrator")
		annotations := namespace.GetAnnotations()
		if annotations == nil {
			common.ToRequestParamsError(g,
				notResourceAllocatedError,
			)
			return
		}
		limits, ok := annotations[constraint.NamespaceAnnotationForNodeResource]
		if !ok {
			common.ToRequestParamsError(g,
				notResourceAllocatedError,
			)
			return
		}

		cds := make(nuwav1.Coordinates, 0)
		err = json.Unmarshal([]byte(limits), &cds)
		if err != nil {
			common.ToRequestParamsError(g, err)
			return
		}
		replicas, err := strconv.ParseUint(deployTemplate.Replicas, 10, 32)
		if err != nil {
			common.ToRequestParamsError(g, err)
			return
		}
		rs := int32(replicas)
		_, cgs := groupBy(cds, rs)

		runtimeObj = &nuwav1.Stone{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Stone",
				APIVersion: "nuwa.nip.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      deployTemplate.AppName,
				Namespace: deployTemplate.Namespace.Value,
				Labels: map[string]string{
					"app": deployTemplate.AppName,
				},
			},
			Spec: nuwav1.StoneSpec{
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Name: deployTemplate.AppName,
						Labels: map[string]string{
							"app": deployTemplate.AppName,
						},
					},
					Spec: corev1.PodSpec{
						Containers: toPodContainer(deployItems...),
					},
				},
				Strategy:    "Release", // TODO
				Coordinates: cgs,
				Service: corev1.ServiceSpec{
					Type: corev1.ServiceType("ClusterIP"),
				},
			},
		}
		runtimeClassGVR = types.ResourceStone
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(runtimeObj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	unstructuredData := &unstructured.Unstructured{Object: unstructuredObj}

	w.generic.SetGroupVersionResource(runtimeClassGVR)
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
