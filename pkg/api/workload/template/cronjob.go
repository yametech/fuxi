package template

import (
	"k8s.io/api/batch/v1beta1"
)

type AttachCronJobRequest struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Container string `json:"container"`
	Shell     string `json:"shell"`
}

type CronJobRequest struct {
	CommonTemplate
	Spec v1beta1.CronJobSpec `json:"spec"`
}

type CronJobResponse struct {
	CronJobList v1beta1.CronJob `json:"cronJob_list"`
}
