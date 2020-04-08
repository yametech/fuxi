package handler

import (
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
)

// WorkloadsAPI all resource operate
type WorkloadsAPI struct {
	deployments *workloadservice.Deployment
	cronJob     *workloadservice.CronJob
	statefulSet *workloadservice.StatefulSet
	daemonSet   *workloadservice.DaemonSet
}

func NewWorkladAPI() *WorkloadsAPI {
	return &WorkloadsAPI{
		deployments: workloadservice.NewDeployment(),
		cronJob:     workloadservice.NewCronJob(),
		statefulSet: workloadservice.NewStatefulSet(),
		daemonSet:   workloadservice.NewDaemonSet(),
	}
}
