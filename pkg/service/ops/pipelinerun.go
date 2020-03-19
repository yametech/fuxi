package ops

import (
	"errors"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/yametech/fuxi/pkg/logging"
	kubeerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

type PipelineRunService interface {
	CreateOrUpdatePipelineRun(pr *PipelineRun) error
	GetPipelineRun(name, namespace string) (*PipelineRun, error)
	PipelineRunList(namespace string) ([]PipelineRun, error)
	PipelineRunDelete(namespace, name string) error
	ReRunPipeline(name, namespace string) error
}

type PipelineRun struct {
	Name             string
	Namespace        string
	Labels           map[string]string
	PipelineRef      *v1alpha1.PipelineRef
	Timeout          *metav1.Duration
	PipelineResource []v1alpha1.PipelineResourceBinding
}

//toPipelineRun map PipelineRun resource
func toPipelineRun(pr *PipelineRun) *v1alpha1.PipelineRun {
	return &v1alpha1.PipelineRun{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:   pr.Name,
			Labels: pr.Labels,
		},
		Spec: v1alpha1.PipelineRunSpec{
			PipelineRef:         pr.PipelineRef,
			PipelineSpec:        nil,
			Resources:           pr.PipelineResource,
			Params:              nil,
			ServiceAccountName:  "default",
			ServiceAccountNames: nil,
			Status:              "",
			Timeout:             pr.Timeout,
			PodTemplate:         v1alpha1.PodTemplate{},
		},
		Status: v1alpha1.PipelineRunStatus{},
	}
}

//CreateOrUpdatePipelineRun create or update a pipeline run
func (ops *Ops) CreateOrUpdatePipelineRun(pr *PipelineRun) error {
	if nil == pr {
		return errors.New("PipelineRun should not be empty")
	}
	if pr.Name == "" && pr.Namespace == "" {
		return errors.New("name or namespace should not be empty")
	}
	oldPipelineRun, err := ops.client.
		PipelineRuns(pr.Namespace).Get(pr.Name, metav1.GetOptions{})

	if err != nil {
		if kubeerr.IsNotFound(err) {
			_, err := ops.client.
				PipelineRuns(pr.Namespace).
				Create(toPipelineRun(pr))
			if err != nil {
				return err
			}
		} else {
			return err
		}

	}

	if oldPipelineRun.Spec.DeepCopy().PipelineSpec.Tasks != nil {
		newPipelineRun := toPipelineRun(pr)
		if !reflect.DeepEqual(oldPipelineRun.Spec, newPipelineRun.Spec) {
			oldPipelineRun.Spec = newPipelineRun.Spec
			_, err = ops.client.PipelineRuns(pr.Namespace).Update(oldPipelineRun)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//GetPipelineRun get a pipeline run Resources
func (ops *Ops) GetPipelineRun(name, namespace string) (*PipelineRun, error) {

	if name == "" && namespace == "" {
		return nil, errors.New("name or namespace should not be empty")
	}
	pr, err := ops.client.
		PipelineRuns(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &PipelineRun{
		Name:             pr.Name,
		Namespace:        pr.Namespace,
		Labels:           pr.Labels,
		PipelineRef:      pr.Spec.PipelineRef,
		Timeout:          pr.Spec.Timeout,
		PipelineResource: pr.Spec.Resources,
	}, nil
}

//PipelineRunList get all pipeline run resources
func (ops *Ops) PipelineRunList(namespace string) ([]PipelineRun, error) {
	if namespace == "" {
		return nil, errors.New("namespace cannot be empty")
	}
	prOpt := metav1.ListOptions{
		LabelSelector: namespace,
	}
	rs, err := ops.client.PipelineRuns(namespace).List(prOpt)
	if err != nil {
		return nil, err
	}
	var prs []PipelineRun
	for _, pr := range rs.Items {
		prs = append(prs, PipelineRun{
			Name:             pr.Name,
			Namespace:        pr.Namespace,
			Labels:           pr.Labels,
			PipelineRef:      pr.Spec.PipelineRef,
			Timeout:          pr.Spec.Timeout,
			PipelineResource: pr.Spec.Resources,
		})
	}

	return prs, nil
}

//PipelineRunDelete delete a pipeline run  resource
func (ops *Ops) PipelineRunDelete(namespace, name string) error {
	if name == "" && namespace == "" {
		return errors.New("name or namespace should not be empty")
	}
	err := ops.client.PipelineRuns(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

//ReRunPipeline rerun a pipeline
func (ops *Ops) ReRunPipeline(name, namespace string) error {

	if name == "" && namespace == "" {
		return errors.New("name or namespace should not be empty")
	}

	pr, err := ops.client.
		PipelineRuns(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	newPipelineRun := pr
	newPipelineRun.Name = ""
	newPipelineRun.Spec.Status = ""
	_ = ops.DeleteTask(name, namespace)

	_, err = ops.client.
		PipelineRuns(pr.Namespace).
		Create(newPipelineRun)
	if err != nil {
		logging.Log.Errorf("an error occurred rerunning the PipelineRun %s in namespace %s: %s", name, namespace, err)
		return err
	}
	return nil
}
