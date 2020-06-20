package ops

import (
	"errors"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	kubeerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TaskService interface {
	CreateOrUpdateTask(task *Task) error
	TaskList(namespace string) ([]Task, error)
	DeleteTask(name, namespace string) error
	GetTask(name, namespace string) (*Task, error)
}

//Task task resource
type Task struct {
	Name      string
	Namespace string
	Inputs    *v1alpha1.Inputs
	Labels    map[string]string
	Outputs   *v1alpha1.Outputs
	Steps     []v1alpha1.Step
}

// toTask map Task resource
func toTask(task *Task) *v1alpha1.Task {
	return &v1alpha1.Task{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:   task.Name,
			Labels: task.Labels,
		},
		Spec: v1alpha1.TaskSpec{
			Inputs:  task.Inputs,
			Outputs: task.Outputs,
			//Steps:        task.Steps,
			//StepTemplate: nil,
			//Sidecars:     nil,
		},
	}
}

//CreateOrUpdateTask creates or updates the task resource
func (ops *Ops) CreateOrUpdateTask(task *Task) error {

	if task == nil {
		return errors.New("the task should not be empty")
	}

	if task.Namespace == "" && task.Name == "" {
		return errors.New("name and namespace should not be empty")
	}

	oldTask, err := ops.client.Tasks(task.Namespace).Get(task.Name, v1.GetOptions{})
	if err != nil {
		if kubeerr.IsNotFound(err) {
			_, err := ops.client.
				Tasks(task.Name).
				Create(toTask(task))
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if oldTask.Name != "" {
		newTask := toTask(task)

		if reflect.DeepEqual(oldTask.Spec, newTask.Spec) {
			oldTask.Spec = newTask.Spec
			_, err = ops.client.Tasks(task.Namespace).Update(oldTask)
			if err != nil {
				return err
			}
		}
	}

	//maybe will return error wrapper? or create the task

	return nil
}

//TaskList Get all task resources with the namespace/labels
func (ops *Ops) TaskList(namespace string) ([]Task, error) {

	if namespace == "" {
		return nil, errors.New("namespace cannot be empty")
	}

	var strVals []string
	strVals = append(strVals, namespace)
	key := "namespace"

	rq, err := labels.NewRequirement(key, selection.Equals, strVals)
	if err != nil {
		return nil, err
	}
	lable := labels.NewSelector().Add(*rq)

	ts, err := ops.informer.Tekton().V1alpha1().Tasks().
		Lister().
		Tasks(namespace).List(lable)

	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, t := range ts {
		tasks = append(tasks, Task{
			Name:      t.Name,
			Namespace: t.Namespace,
			Inputs:    t.Spec.Inputs,
			Labels:    t.Labels,
			Outputs:   t.Spec.Outputs,
			Steps:     t.Spec.Steps,
		})
	}

	return tasks, nil
}

//DeleteTask deletes a task resource
func (ops *Ops) DeleteTask(name, namespace string) error {

	if name == "" && namespace == "" {
		return errors.New("name and namespace should not be empty")
	}

	err := ops.client.Tasks(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}

//GetTask get a task resource
func (ops *Ops) GetTask(name, namespace string) (*Task, error) {

	if name == "" && namespace == "" {
		return nil, errors.New("name and namespace should not be empty")
	}

	task, err := ops.client.Tasks(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &Task{
		Name:      task.Name,
		Namespace: task.Namespace,
		Inputs:    task.Spec.Inputs,
		Labels:    task.Labels,
		Outputs:   task.Spec.Outputs,
		Steps:     task.Spec.Steps,
	}, nil
}
