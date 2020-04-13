package handler

import (
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
)

// WorkloadsAPI all resource operate
type WorkloadsAPI struct {
	deployments   *workloadservice.Deployment
	job           *workloadservice.Job
	cronJob       *workloadservice.CronJob
	statefulSet   *workloadservice.StatefulSet
	daemonSet     *workloadservice.DaemonSet
	replicaSet    *workloadservice.ReplicaSet
	pod           *workloadservice.Pod
	event         *workloadservice.Event
	node          *workloadservice.Node
	configMaps    *workloadservice.ConfigMaps
	secret        *workloadservice.Secrets
	resourceQuota *workloadservice.ResourceQuota
}

func NewWorkladAPI() *WorkloadsAPI {
	return &WorkloadsAPI{
		deployments:   workloadservice.NewDeployment(),
		cronJob:       workloadservice.NewCronJob(),
		statefulSet:   workloadservice.NewStatefulSet(),
		daemonSet:     workloadservice.NewDaemonSet(),
		job:           workloadservice.NewJob(),
		replicaSet:    workloadservice.NewReplicaSet(),
		pod:           workloadservice.NewPod(),
		event:         workloadservice.NewEvent(),
		node:          workloadservice.NewNode(),
		configMaps:    workloadservice.NewConfigMaps(),
		secret:        workloadservice.NewSecrets(),
		resourceQuota: workloadservice.NewResourceQuota(),
	}
}
