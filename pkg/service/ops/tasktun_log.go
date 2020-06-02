// Copyright © 2019 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ops

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/yametech/fuxi/pkg/service/common"
	"github.com/yametech/fuxi/pkg/service/ops/pods"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"knative.dev/pkg/apis/duck/v1beta1"
)

const (
	MsgTRNotFoundErr = "Unable to get Taskrun"
)

type step struct {
	name      string
	container string
	state     corev1.ContainerState
}

func (s *step) hasStarted() bool {
	return s.state.Waiting == nil
}

func (log *Logger) readTaskLog(name, ns string) (<-chan Log, <-chan error, error) {
	tr, err := TektonClient.TektonV1alpha1().TaskRuns(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %s", MsgTRNotFoundErr, err)
	}

	log.formTaskName(tr)

	//if r.follow {
	//
	//}
	//return r.readAvailableTaskLogs(tr)
	return log.readLiveTaskLogs()
}

func (log *Logger) formTaskName(tr *v1alpha1.TaskRun) {
	if log.task != "" {
		return
	}

	if name, ok := tr.Labels["tekton.dev/pipelineTask"]; ok {
		log.task = name
		return
	}

	if tr.Spec.TaskRef != nil {
		log.task = tr.Spec.TaskRef.Name
		return
	}

	log.task = fmt.Sprintf("Task %d", log.number)
}

func (log *Logger) readLiveTaskLogs() (<-chan Log, <-chan error, error) {
	tr, err := log.waitUntilTaskPodNameAvailable(10)
	if err != nil {
		return nil, nil, err
	}

	var (
		podName = tr.Status.PodName
		kube    = common.SharedK8sClient.ClientV1
	)
	//New -> NewWithDefaults
	p := pods.NewWithDefaults(podName, log.ns, kube)
	pod, err := p.Wait()
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("task %s failed: %s. Run tkn tr desc %s for more details.", log.task, strings.TrimSpace(err.Error()), tr.Name))
	}

	steps := filterSteps(pod, true, []string{})
	logC, errC := log.readStepsLogs(steps, p, true)
	return logC, errC, err
}

func (log *Logger) readStepsLogs(steps []*step, pod *pods.Pod, follow bool) (<-chan Log, <-chan error) {
	logC := make(chan Log)
	errC := make(chan error)

	go func() {
		defer close(logC)
		defer close(errC)

		for _, step := range steps {
			if !follow && !step.hasStarted() {
				continue
			}

			container := pod.Container(step.container)
			podC, perrC, err := container.LogReader(follow).Read()
			if err != nil {
				errC <- fmt.Errorf("error in getting logs for step %s: %s", step.name, err)
				continue
			}

			for podC != nil || perrC != nil {
				select {
				case l, ok := <-podC:
					if !ok {
						podC = nil
						logC <- Log{Task: log.task, Step: step.name, Log: "EOFLOG"}
						continue
					}
					logC <- Log{Task: log.task, Step: step.name, Log: l.Log}

				case e, ok := <-perrC:
					if !ok {
						perrC = nil
						continue
					}

					errC <- fmt.Errorf("failed to get logs for %s: %s", step.name, e)
				}
			}

			if err := container.Status(); err != nil {
				errC <- err
				return
			}
		}
	}()

	return logC, errC
}

// Reading of logs should wait until the name of the pod is
// updated in the status. Open a watch channel on the task run
// and keep checking the status until the pod name updates
// or the timeout is reached.
func (log *Logger) waitUntilTaskPodNameAvailable(timeout time.Duration) (*v1alpha1.TaskRun, error) {
	var first = true
	opts := metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", log.run).String(),
	}

	run, err := TektonClient.TektonV1alpha1().TaskRuns(log.ns).Get(log.run, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if run.Status.PodName != "" {
		return run, nil
	}

	watchRun, err := TektonClient.TektonV1alpha1().TaskRuns(log.ns).Watch(opts)
	if err != nil {
		return nil, err
	}
	for {
		select {
		case event := <-watchRun.ResultChan():
			run := event.Object.(*v1alpha1.TaskRun)
			if run.Status.PodName != "" {
				watchRun.Stop()
				return run, nil
			}
			if first {
				first = false
			}
		case <-time.After(timeout * time.Second):
			watchRun.Stop()

			// Check if taskrun failed on start up
			if err = hasTaskRunFailed(run.Status.Conditions, log.task); err != nil {
				return nil, err
			}

			return nil, fmt.Errorf("task %s create has not started yet or pod for task not yet available", log.task)
		}
	}
}

func filterSteps(pod *corev1.Pod, allSteps bool, stepsGiven []string) []*step {
	steps := []*step{}
	stepsInPod := getSteps(pod)

	if allSteps {
		steps = append(steps, getInitSteps(pod)...)
	}

	if len(stepsGiven) == 0 {
		steps = append(steps, stepsInPod...)
		return steps
	}

	stepsToAdd := map[string]bool{}
	for _, s := range stepsGiven {
		stepsToAdd[s] = true
	}

	for _, sp := range stepsInPod {
		if stepsToAdd[sp.name] {
			steps = append(steps, sp)
		}
	}

	return steps
}

func getInitSteps(pod *corev1.Pod) []*step {
	status := map[string]corev1.ContainerState{}
	for _, ics := range pod.Status.InitContainerStatuses {
		status[ics.Name] = ics.State
	}

	steps := []*step{}
	for _, ic := range pod.Spec.InitContainers {
		steps = append(steps, &step{
			name:      strings.TrimPrefix(ic.Name, "step-"),
			container: ic.Name,
			state:     status[ic.Name],
		})
	}

	return steps
}

func getSteps(pod *corev1.Pod) []*step {
	status := map[string]corev1.ContainerState{}
	for _, cs := range pod.Status.ContainerStatuses {
		status[cs.Name] = cs.State
	}

	steps := []*step{}
	for _, c := range pod.Spec.Containers {
		steps = append(steps, &step{
			name:      strings.TrimPrefix(c.Name, "step-"),
			container: c.Name,
			state:     status[c.Name],
		})
	}

	return steps
}

func hasTaskRunFailed(trConditions v1beta1.Conditions, taskName string) error {
	if len(trConditions) != 0 && trConditions[0].Status == corev1.ConditionFalse {
		return fmt.Errorf("task %s has failed: %s", taskName, trConditions[0].Message)
	}
	return nil
}
