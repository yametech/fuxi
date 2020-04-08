package handler

import (
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
)

// WorkloadsAPI all resource operate
type WorkloadsAPI struct {
	deployments *workloadservice.Deployment
}

func NewWorkladAPI() *WorkloadsAPI {
	return &WorkloadsAPI{
		deployments: workloadservice.NewDeployment(),
	}
}
