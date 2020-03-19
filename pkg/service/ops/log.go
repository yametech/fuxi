package ops

import (
	"bytes"
	"errors"
	"github.com/yametech/fuxi/pkg/k8s/client"
	"github.com/yametech/fuxi/pkg/logging"

	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type LogService interface {
	GetTaskRunLog(name, namespace string) (*TaskRunLog, error)
	GetPipelineRunLog(name, namespace string) (PipelineRunLog, error)
}

type PipelineRunLog []TaskRunLog

//TaskRunLog
type TaskRunLog struct {
	PodName        string
	StepContainers []LogContainer
	PodContainers  []LogContainer
	InitContainers []LogContainer
}

//LogContainer
type LogContainer struct {
	Name string
	Logs []string
}

const ContainerPrefix = "build-step-"

//GetTaskRunLog  Get the logs for a given task run by name in a given namespace
func (ops *Ops) GetTaskRunLog(name, namespace string) (*TaskRunLog, error) {

	logging.Log.Debugf("In getTaskRunLog - name: %s, namespace: %s", name, namespace)
	taskRunsInterface := ops.client.TaskRuns(namespace)
	taskRun, err := taskRunsInterface.Get(name, metav1.GetOptions{})
	if err != nil || taskRun.Status.PodName == "" {
		return nil, errors.New("cloud not get task run")
	}

	pod, err := client.K8sClient.CoreV1().Pods(namespace).Get(taskRun.Status.PodName, metav1.GetOptions{})
	if err != nil {
		return nil, errors.New("cloud not get pod")
	}
	return ops.makeTaskRunLog(pod), nil
}

func (ops *Ops) makeTaskRunLog(pod *v1.Pod) *TaskRunLog {
	podContainers := pod.Spec.Containers
	initContainers := pod.Spec.InitContainers

	taskRunLog := &TaskRunLog{PodName: pod.Name}
	buf := new(bytes.Buffer)
	setContainers := func(containers []v1.Container, filter func(l LogContainer)) {
		for _, container := range containers {
			buf.Reset()
			step := LogContainer{Name: container.Name}
			req := client.K8sClient.CoreV1().Pods(pod.GetNamespace()).GetLogs(pod.Name, &v1.PodLogOptions{Container: container.Name})
			if req.URL().Path == "" {
				continue
			}
			podLogs, _ := req.Stream()
			if podLogs == nil {
				continue
			}
			_, err := io.Copy(buf, podLogs)
			if err != nil {
				podLogs.Close()
				continue
			}
			logs := strings.Split(buf.String(), "\n")
			for _, log := range logs {
				if log != "" {
					step.Logs = append(step.Logs, log)
				}
			}
			filter(step)
			podLogs.Close()
		}
	}
	setContainers(initContainers, func(l LogContainer) {
		taskRunLog.InitContainers = append(taskRunLog.InitContainers, l)
	})
	setContainers(podContainers, func(l LogContainer) {
		if strings.HasPrefix(l.Name, ContainerPrefix) {
			taskRunLog.StepContainers = append(taskRunLog.StepContainers, l)
		} else {
			taskRunLog.PodContainers = append(taskRunLog.PodContainers, l)
		}
	})
	return taskRunLog
}

//GetPipelineRunLog  Get the logs for a given pipelinerun by name in a given namespace
func (ops *Ops) GetPipelineRunLog(name, namespace string) (PipelineRunLog, error) {

	pipelinerun, err := ops.client.PipelineRuns(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, errors.New("cloud not get pipelinerun")
	}

	var pipelineRunLogs PipelineRunLog
	for _, taskrunstatus := range pipelinerun.Status.TaskRuns {
		podname := taskrunstatus.Status.PodName
		pod, err := client.K8sClient.CoreV1().Pods(namespace).Get(podname, metav1.GetOptions{})
		if err != nil {
			continue
		}
		pipelineRunLogs = append(pipelineRunLogs, *ops.makeTaskRunLog(pod))
	}

	return pipelineRunLogs, nil
}
