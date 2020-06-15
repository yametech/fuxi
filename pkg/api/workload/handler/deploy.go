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
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/http"
	"strconv"
)

type deployTemplate struct {
	AppName   string `json:"appName"`
	Namespace struct {
		Value string `json:"value"`
	} `json:"namespace"`
	Replicas     string `json:"replicas"`
	TemplateName string `json:"templateName"`
}

func workloadsTemplateToServiceSpec(wt *workloadsTemplate) (*corev1.ServiceSpec, error) {
	serviceSpec := &corev1.ServiceSpec{
		Type: corev1.ServiceType(wt.Service.Type),
	}
	for _, item := range wt.Service.Ports {
		port, err := strconv.ParseInt(item.Port, 10, 32)
		if err != nil {
			return nil, err
		}
		newItem := corev1.ServicePort{
			Name:       item.Name,
			Protocol:   corev1.Protocol(item.Protocol),
			Port:       int32(port),
			TargetPort: intstr.Parse(item.TargetPort),
		}
		serviceSpec.Ports = append(serviceSpec.Ports, newItem)
	}
	return serviceSpec, nil
}

func string2int32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int32(i)
}

func workloadsTemplateToPodContainers(wt *workloadsTemplate) []corev1.Container {
	containers := make([]corev1.Container, 0)
	for _, item := range wt.Metadata {
		volumeMounts := make([]corev1.VolumeMount, 0)
		for _, subItem := range item.VolumeMounts.Items {
			volumeMounts = append(volumeMounts,
				corev1.VolumeMount{
					Name:      subItem.Name,
					MountPath: subItem.MountPath,
				})
		}
		container := corev1.Container{
			Name:            item.Base.Name,
			Image:           item.Base.Image,
			ImagePullPolicy: corev1.PullPolicy(item.Base.ImagePullPolicy),
			Resources:
			corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					corev1.ResourceName("cpu"): resource.MustParse(
						fmt.Sprintf("%s", item.Base.Resource.Limits.CPU),
					),
					corev1.ResourceName("memory"): resource.MustParse(
						fmt.Sprintf("%sM", item.Base.Resource.Limits.Memory),
					),
				},
				Requests: corev1.ResourceList{
					corev1.ResourceName("cpu"): resource.MustParse(
						fmt.Sprintf("%s", item.Base.Resource.Requests.CPU)),
					corev1.ResourceName("memory"): resource.MustParse(
						fmt.Sprintf("%sM", item.Base.Resource.Requests.Memory),
					),
				},
			},
			VolumeMounts: volumeMounts,
		}
		// Command
		for _, cmd := range item.Commands {
			container.Command = append(container.Command, cmd)
		}
		// Args
		for _, arg := range item.Args {
			container.Args = append(container.Command, arg)
		}
		// Environment TODO
		envs := make([]corev1.EnvVar, 0)
		for _, evnConfig := range item.Environment {
			switch evnConfig.Type {
			case "Configuration":
				//env := &corev1.EnvVar{
				//	Name: evnConfig.OneEnvConfig.Name,
				//	ValueFrom: &corev1.EnvVarSource{
				//		ConfigMapKeyRef:
				//		&corev1.ConfigMapEnvSource{
				//
				//		},
				//	},
				//}
			case "Secret":
			case "Other":
			default:
			}
		}
		container.Env = envs

		// readinessProbe
		//container.ReadinessProbe = &corev1.Probe{
		//	Handler: corev1.Handler{
		//		Exec: &corev1.ExecAction{
		//			Command: []string{item.ReadyProbe.Pattern.Command},
		//		},
		//		HTTPGet: &corev1.HTTPGetAction{
		//			Path: item.ReadyProbe.Pattern.URL,
		//			Port: intstr.Parse(item.ReadyProbe.Pattern.HTTPPort),
		//		},
		//		TCPSocket: &corev1.TCPSocketAction{
		//			Port: intstr.Parse(item.ReadyProbe.Pattern.TCPPort),
		//		},
		//	},
		//	InitialDelaySeconds: string2int32(item.ReadyProbe.Delay),
		//	TimeoutSeconds:      string2int32(item.ReadyProbe.Timeout),
		//	PeriodSeconds:       string2int32(item.ReadyProbe.Cycle),
		//	FailureThreshold:    string2int32(item.ReadyProbe.RetryCount),
		//}
		//
		//// livenessProbe
		//container.LivenessProbe = &corev1.Probe{
		//	Handler: corev1.Handler{
		//		Exec: &corev1.ExecAction{
		//			Command: []string{item.LiveProbe.Pattern.Command},
		//		},
		//		HTTPGet: &corev1.HTTPGetAction{
		//			Path: item.ReadyProbe.Pattern.URL,
		//			Port: intstr.Parse(item.LiveProbe.Pattern.HTTPPort),
		//		},
		//		TCPSocket: &corev1.TCPSocketAction{
		//			Port: intstr.Parse(item.LiveProbe.Pattern.TCPPort),
		//		},
		//	},
		//	InitialDelaySeconds: string2int32(item.LiveProbe.Delay),
		//	TimeoutSeconds:      string2int32(item.LiveProbe.Timeout),
		//	PeriodSeconds:       string2int32(item.LiveProbe.Cycle),
		//	FailureThreshold:    string2int32(item.LiveProbe.RetryCount),
		//}
		//// LifeCycle
		//container.Lifecycle = &corev1.Lifecycle{
		//	PostStart: &corev1.Handler{
		//		Exec: &corev1.ExecAction{
		//			Command: []string{item.LifeCycle.PostStart.Command},
		//		},
		//		HTTPGet: &corev1.HTTPGetAction{
		//			Path: item.LifeCycle.PostStart.URL,
		//			Port: intstr.Parse(item.LifeCycle.PostStart.HTTPPort),
		//		},
		//		TCPSocket: &corev1.TCPSocketAction{
		//			Port: intstr.Parse(item.LifeCycle.PostStart.TCPPort),
		//		},
		//	},
		//	PreStop: &corev1.Handler{
		//		Exec: &corev1.ExecAction{
		//			Command: []string{item.LifeCycle.PreStop.Command},
		//		},
		//		HTTPGet: &corev1.HTTPGetAction{
		//			Path: item.ReadyProbe.Pattern.URL,
		//			Port: intstr.Parse(item.LifeCycle.PreStop.HTTPPort),
		//		},
		//		TCPSocket: &corev1.TCPSocketAction{
		//			Port: intstr.Parse(item.LifeCycle.PreStop.TCPPort),
		//		},
		//	},
		//}

		containers = append(containers, container)
	}
	return containers
}

func workloadsTemplateToVolumeClaims(wt *workloadsTemplate) []corev1.PersistentVolumeClaim {
	pvcs := make([]corev1.PersistentVolumeClaim, 0)
	for _, item := range wt.VolumeClaims {
		pvc := corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name: item.Metadata.Name,
			},
			Spec: corev1.PersistentVolumeClaimSpec{},
		}
		accessModes := make([]corev1.PersistentVolumeAccessMode, 0)
		for _, subItem := range item.Spec.AccessModes {
			accessModes = append(accessModes, corev1.PersistentVolumeAccessMode(subItem))
		}
		resourceRequire := corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceName("storage"): resource.MustParse(item.Spec.Resources.Requests.Storage),
			},
		}
		if item.Metadata.IsUseDefaultStorageClass {
			pvc.ObjectMeta.Annotations = map[string]string{
				"volume.alpha.kubernetes.io/storage-class": "default",
			}
			pvc.Spec = corev1.PersistentVolumeClaimSpec{
				AccessModes: accessModes,
				Resources:   resourceRequire,
			}
		} else {
			pvc.Spec = corev1.PersistentVolumeClaimSpec{
				StorageClassName: &item.Spec.StorageClassName,
				AccessModes:      accessModes,
				Resources:        resourceRequire,
			}

		}
		pvcs = append(pvcs, pvc)
	}
	return pvcs
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

func workloadsTemplateImagePullSecrets(w *workloadsTemplate) []corev1.LocalObjectReference {
	result := make([]corev1.LocalObjectReference, 0)
	for _, item := range w.Metadata {
		result = append(result, corev1.LocalObjectReference{Name: item.Base.ImagePullSecret})
	}
	return result
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

	workloadsTemplate := &workloadsTemplate{
		Metadata:     make(metadataTemplate, 0),
		Service:      serviceTemplate{},
		VolumeClaims: make(volumeClaimsTemplate, 0),
	}
	if err := json.Unmarshal([]byte(*workloads.Spec.Metadata), &workloadsTemplate.Metadata); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	if err := json.Unmarshal([]byte(*workloads.Spec.Service), &workloadsTemplate.Service); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	if err := json.Unmarshal([]byte(workloads.Spec.VolumeClaims), &workloadsTemplate.VolumeClaims); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	var runtimeObj runtime.Object

	var runtimeClassGVR schema.GroupVersionResource
	switch *workloads.Spec.ResourceType {
	case "Stone":
		obj, err := w.namespace.Get("", deployTemplate.Namespace.Value)
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

		serviceTemplate, err := workloadsTemplateToServiceSpec(workloadsTemplate)
		if err != nil {
			if err != nil {
				common.ToRequestParamsError(g, err)
				return
			}
		}

		// PodSpec
		podSpec := corev1.PodSpec{}
		podSpec.Containers = workloadsTemplateToPodContainers(workloadsTemplate)
		podSpec.ImagePullSecrets = workloadsTemplateImagePullSecrets(workloadsTemplate)

		// Labels
		labels := map[string]string{
			"app":               deployTemplate.AppName,
			"app-template-name": deployTemplate.TemplateName,
		}
		runtimeObj = &nuwav1.Stone{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Stone",
				APIVersion: "nuwa.nip.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      deployTemplate.AppName,
				Namespace: deployTemplate.Namespace.Value,
				Labels:    labels,
			},
			Spec: nuwav1.StoneSpec{
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Name:   deployTemplate.AppName,
						Labels: labels,
					},
					Spec: podSpec,
				},
				Strategy:             "Release", // TODO
				Coordinates:          cgs,
				Service:              *serviceTemplate,
				VolumeClaimTemplates: workloadsTemplateToVolumeClaims(workloadsTemplate),
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