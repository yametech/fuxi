package ops

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/yametech/fuxi/pkg/k8s/client"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/tekton"
	corev1 "k8s.io/api/core/v1"
	"sync"

	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type LogService interface {
	GetTaskRunLog(name, namespace string) (*TaskRunLog, error)
	GetPipelineRunLog(name, namespace string) (PipelineRunLog, error)
	GetTaskRealLog(name, namespace string,logs chan []string) error
    ReadLivePipelineLogs(name, namespace string,tasks []string) (<-chan Log, <-chan error, error)
}

// Log represents data to write on log channel
type Log struct {
	Pipeline string
	Task     string
	Step     string
	Log      string
}


type Logger struct {
	number int
	run string
	task string
	ns string
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


func (ops *Ops)GetTaskRealLog(name, namespace string,logs chan []string) error {

	pipelinerun, err := ops.client.PipelineRuns(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	for _, taskrunstatus := range pipelinerun.Status.TaskRuns {
		podname := taskrunstatus.Status.PodName
		pod, err := client.K8sClient.CoreV1().Pods(namespace).Get(podname, metav1.GetOptions{})
		if err != nil {
			continue
		}

		buf := new(bytes.Buffer)
		if pod.Spec.Containers != nil {
			for _, container := range pod.Spec.Containers {
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
				logs <- strings.Split(buf.String(), "\n")
			}
		}
	}

	return nil

}


func (ops *Ops) ReadLivePipelineLogs(name, namespace string,tasks []string) (<-chan Log, <-chan error, error) {
	pr, err := ops.client.
		PipelineRuns(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil,nil, err
	}
	logC := make(chan Log)
	errC := make(chan error)

	go func() {
		defer close(logC)
		defer close(errC)

		prTracker := NewTracker(name, namespace, tekton.TektonClient)
		if tasks == nil {
			tasks = make([]string,0)
		}
		trC := prTracker.Monitor(tasks)

		wg := sync.WaitGroup{}
		taskIndex := 0

		for trs := range trC {
			wg.Add(len(trs))

			for _, run := range trs {
				taskIndex++
				// NOTE: passing tr, taskIdx to avoid data race
				go func(tr Run, taskNum int) {
					defer wg.Done()

					// clone the object to keep task number and name separately
					c := ops.clone()
					c.setUpTask(taskNum, tr)
					c.pipeLogs(logC, errC)
				}(run, taskIndex)
			}
		}

		wg.Wait()

		if !empty(pr.Status) && pr.Status.Conditions[0].Status == corev1.ConditionFalse {
			errC <- fmt.Errorf(pr.Status.Conditions[0].Message)
		}
	}()

	return logC, errC, nil
}


func (log *Logger) setUpTask(taskNumber int, tr Run) {
	log.setNumber(taskNumber)
	log.setRun(tr.Name)
	log.setTask(tr.Task)
}

func (o *Ops)clone() *Logger{
	l := *o.log
	return &l
}

func (log *Logger) setNumber(number int) {
	log.number = number
}

func (log *Logger) setRun(run string) {
	log.run = run
}

func (log *Logger) setTask(task string) {
	log.task = task
}


func empty(status v1alpha1.PipelineRunStatus) bool {
	if status.Conditions == nil {
		return true
	}
	return len(status.Conditions) == 0
}



func (log *Logger) pipeLogs(logC chan<- Log, errC chan<- error) {
	tlogC, terrC, err := log.readTaskLog(log.task,log.ns)
	if err != nil {
		errC <- err
		return
	}

	for tlogC != nil || terrC != nil {
		select {
		case l, ok := <-tlogC:
			if !ok {
				tlogC = nil
				continue
			}
			logC <- Log{Task: l.Task, Step: l.Step, Log: l.Log}

		case e, ok := <-terrC:
			if !ok {
				terrC = nil
				continue
			}
			errC <- fmt.Errorf("failed to get logs for task %s : %s", log.task, e)
		}
	}
}
