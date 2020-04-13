package ops

import (
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"reflect"
)

type PipelineService interface {
	CreateOrUpdatePipeline(pipeLine Pipeline) error
	PipelineList(namespace string) ([]Pipeline, error)
	PipelineDelete(namespace, name string) error
	GetPipeline(namespace, name string) (*Pipeline, error)
	CancelPipelineRun(namespace, name string)  error
}

type Pipeline struct {
	Name       string
	Namespace  string
	Labels     map[string]string
	Tasks      []v1alpha1.PipelineTask
	Resources  []v1alpha1.PipelineDeclaredResource
	ParamSpecs []v1alpha1.ParamSpec
}

//toPipelineResource map Pipeline  resource
func toPipeline(name string,
	labels map[string]string,
	resources []v1alpha1.PipelineDeclaredResource,
	tasks []v1alpha1.PipelineTask,
	paramSpecs []v1alpha1.ParamSpec) *v1alpha1.Pipeline {
	return &v1alpha1.Pipeline{
		TypeMeta: v1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
		Spec: v1alpha1.PipelineSpec{
			Resources: resources,
			Tasks:     tasks,
			Params:    paramSpecs,
		},
		Status: v1alpha1.PipelineStatus{},
	}
}


func (ops *Ops) CreateOrUpdatePipeline(pipeLine Pipeline) error {

	oldPipeline, err := ops.client.
		Pipelines(pipeLine.Namespace).Get(pipeLine.Name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err := ops.client.
				Pipelines(pipeLine.Namespace).
				Create(toPipeline(
					pipeLine.Name,
					pipeLine.Labels,
					pipeLine.Resources,
					pipeLine.Tasks,
					pipeLine.ParamSpecs))
			if err != nil {
				return err
			}
		} else {
			return err
		}

	}

	if oldPipeline.Spec.DeepCopy().Tasks != nil {
		newPipeline := toPipeline(
			pipeLine.Name,
			pipeLine.Labels,
			pipeLine.Resources,
			pipeLine.Tasks,
			pipeLine.ParamSpecs)

		if !reflect.DeepEqual(oldPipeline.Spec, newPipeline.Spec) {
			oldPipeline.Spec = newPipeline.Spec
			_, err = ops.client.Pipelines(pipeLine.Namespace).Update(oldPipeline)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//PipelineList Get all  Pipeline resources with the namespace/labels
func (ops *Ops) PipelineList(namespace string) ([]Pipeline, error) {

	var strVals []string
	strVals = append(strVals,namespace)
	key := "namespace"
	rq := labels.Requirement{key,selection.Equals,strVals}
	lable := labels.NewSelector().Add(rq)

	ps, err :=ops.informer.Tekton().V1alpha1().
		Pipelines().
		Lister().
		Pipelines("").List(lable)

	if err != nil {
		return nil, err
	}

	var pipelines []Pipeline
	for _, p := range ps {
		pipelines = append(pipelines, Pipeline{
			Name:       p.Name,
			Namespace:  p.Namespace,
			Labels:     p.Labels,
			Tasks:      p.Spec.Tasks,
			Resources:  p.Spec.Resources,
			ParamSpecs: p.Spec.Params,
		})
	}

	return pipelines, nil
}

//PipelineDelete delete a pipeline resource
func (ops *Ops) PipelineDelete(namespace, name string) error {
	err := ops.client.Pipelines(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

//PipelineDelete delete a pipeline resource
func (ops *Ops) GetPipeline(namespace, name string) (*Pipeline, error) {
	p, err := ops.client.Pipelines(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return &Pipeline{
		Name:       p.Name,
		Namespace:  p.Namespace,
		Labels:     p.Labels,
		Tasks:      p.Spec.Tasks,
		Resources:  p.Spec.Resources,
		ParamSpecs: p.Spec.Params,
	}, nil
}

func (ops *Ops)CancelPipelineRun(namespace, name string)  error{

	pr, err :=ops.client.PipelineRuns(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to find pipelinerun: %s", name)
	}

	status := pr.Status.Conditions[0].Status
	//ConditionFalse == Failed
	//ConditionTrue == Succeeded
	//ConditionUnknown == Running
	if status == corev1.ConditionFalse || status==corev1.ConditionTrue || status==corev1.ConditionUnknown {
		return fmt.Errorf("failed to cancel pipelinerun %s: pipelinerun has already finished execution", name)
	}

	pr.Spec.Status = v1alpha1.PipelineRunSpecStatusCancelled
	_, err =ops.client.PipelineRuns(namespace).Update(pr)
	if err != nil {
		return fmt.Errorf("failed to cancel pipelinerun: %s, err: %s", name, err.Error())
	}

	return nil
}