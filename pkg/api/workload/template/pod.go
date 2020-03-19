package template

import corev1 "k8s.io/api/core/v1"

type PodRequest struct {
	Model CommonTemplate `json:"model" from:"model"`
	Spec  corev1.PodSpec `json:"spec"`
}

type PodResponse struct {
	PodList corev1.PodList `json:"pod_list"`
}
