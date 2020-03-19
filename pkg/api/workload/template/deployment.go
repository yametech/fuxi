package template

import appsv1 "k8s.io/api/apps/v1"

type DeploymentRequest struct {
	Model CommonTemplate        `json:"model" from:"model"`
	Spec  appsv1.DeploymentSpec `json:"spec"`
}
