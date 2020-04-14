package handler

import (
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
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
	}
}
