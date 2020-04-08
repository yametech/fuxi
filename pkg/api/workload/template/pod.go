package template

import corev1 "k8s.io/api/core/v1"

type AttachPodRequest struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Container string `json:"container"`
	Shell     string `json:"shell"`
}

type PodRequest struct {
	CommonTemplate
	Spec corev1.PodSpec `json:"spec"`
}

type PodResponse struct {
	PodList corev1.PodList `json:"pod_list"`
}
