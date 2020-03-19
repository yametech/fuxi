package ops

import (
	"errors"
	"sort"
	"strconv"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/yametech/fuxi/pkg/logging"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type PipelineRunService interface {
	CreateOrUpdatePipelineRun(pr *PipelineRun) error
	GetPipelineRun(name, namespace string) (*PipelineRun, error)
	GetPipelineRunList(namespace string, labelsMatch map[string]string) ([]PipelineRun, error)
	PipelineRunDelete(namespace, name string) error
	ReRunPipeline(name, namespace string) error
	GetLatestPipelineRunList(namespace string, labels map[string]string) ([]PipelineRun, error)
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

// pipelinerun name: xxx_1,xxx_2,xxx_3,xxx_4,xxx_5
// bable:
//      name:xxx
//      namespace:"fjl"
//      number:"1"
//      latest:"true"
//the labels include the pipeline run name
//GetLatestPipelineRunList get all latest pipeline run
func (ops *Ops) GetLatestPipelineRunList(namespace string, labels map[string]string) ([]PipelineRun, error) {

	//list all pipeline run flit by labels(namespace:"",latest:"true"),and then will find all pipeline run
	//in current namespace/department
	prs, err := ops.GetPipelineRunList(namespace, labels)
	if err != nil {
		return nil, err
	}

	newMap := make(map[string]string)
	newLables := make(map[string]map[string]string)
	for _, pr := range prs {

		newMap["name"] = pr.Labels["name"]
		newMap["namespace"] = pr.Labels["namespace"]
		newMap["latest"] = "true"
		newLables["name"] = newMap

	}

	var pipelinerunList []PipelineRun
	for _, lb := range newLables {

		prs, err := ops.GetPipelineRunList(namespace, lb)
		if err != nil {
			return nil, err
		}
		if len(prs) == 0 {
			continue
		}

		var values []int
		prMap := make(map[int]PipelineRun)
		for _, pr := range prs {
			val, err := strconv.ParseInt(pr.Labels["number"], 10, 64)
			if err != nil {
				return nil, errors.New("parse int error" + err.Error())
			}
			v := int(val)
			values = append(values, v)
			prMap[v] = pr

		}
		sort.Ints(values)

		pipelinerunList = append(pipelinerunList, prMap[values[len(values)-1]])
	}

	return pipelinerunList, nil

}

// find all pipeline run by labels,if > 5 ,will delete the first task run and then.
//pr.Labels["name"] = pr.Name
//pr.Labels["namespace"] = pr.Namespace
//labels:
// name: A  name: B  name: C
// number:1  number:2   number:3
// namespace: ns
// latest:"true" / ""
//------------idea:(find min number and then update it)
// name: A-1  name: A-2  name: A-3
// number:4  number:2   number:3
// latest:true
//CreateOrUpdatePipelineRun create or update a pipeline run
func (ops *Ops) CreateOrUpdatePipelineRun(pr *PipelineRun) error {

	if nil == pr {
		return errors.New("PipelineRun should not be empty")
	}
	if pr.Name == "" && pr.Namespace == "" {
		return errors.New("name or namespace should not be empty")
	}

	prs, err := ops.GetPipelineRunList(pr.Namespace, pr.Labels)
	if err != nil {
		return err
	}
	var values []int
	prMap := make(map[int]PipelineRun)
	length := len(prs)
	if length == 0 {

		//pr.Labels["name"] = pr.Name
		//pr.Labels["namespace"] = pr.Namespace
		pr.Labels["latest"] = "true"
		pr.Labels["number"] = "1"
		pr.Name = pr.Name + "-1"

		_, err := ops.client.
			PipelineRuns(pr.Namespace).
			Create(toPipelineRun(pr))
		if err != nil {
			return err
		}

		return nil

	}

	maxHistoryVersionNumber := 5
	//1,2,3,4,5 ----> 6,2,3,4,5
	for _, pr := range prs {
		val, err := strconv.ParseInt(pr.Labels["number"], 10, 64)
		if err != nil {
			return errors.New("parse int error" + err.Error())
		}
		v := int(val)
		values = append(values, v)
		prMap[v] = pr

	}

	//sort it
	sort.Ints(values)

	if length != maxHistoryVersionNumber {

		updateP := prMap[values[len(values)-1]]
		updateP.Labels["latest"] = ""
		_, err := ops.client.PipelineRuns(updateP.Namespace).Update(toPipelineRun(&updateP))
		if err != nil {
			return err
		}

		//++number
		pr.Labels["number"] = string(values[length-1] + 1)
		//and then create the pipeline run resource
		_, err = ops.client.
			PipelineRuns(pr.Namespace).
			Create(toPipelineRun(pr))

		if err != nil {
			return err
		}
		return nil

	}

	p := prMap[values[0]]
	if length == maxHistoryVersionNumber {

		//delete the  fifth resource
		err := ops.PipelineRunDelete(p.Namespace, p.Name)
		if err != nil {
			return err
		}

		//++number
		pr.Labels["number"] = string(values[l-1] + 1)
		//and then create the pipeline run resource
		_, err = ops.client.
			PipelineRuns(pr.Namespace).
			Create(toPipelineRun(pr))

		if err != nil {
			return err
		}
		return nil

	}
	return nil
}

//GetPipelineRun get a pipeline run Resources
func (ops *Ops) GetPipelineRun(name, namespace string) (*PipelineRun, error) {

	if name == "" && namespace == "" {
		return nil, errors.New("name or nam espace should not be empty")
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
func (ops *Ops) GetPipelineRunList(namespace string, labelsMatch map[string]string) ([]PipelineRun, error) {
	if namespace == "" {
		return nil, errors.New("namespace cannot be empty")
	}
	//labelsMatch := make(map[string]string)
	//labelsMatch["a"]= ""
	var labelSelector metav1.LabelSelector
	labelSelector.MatchLabels = labelsMatch
	labelMap, err := metav1.LabelSelectorAsMap(&labelSelector)
	if err != nil {
		return nil, err
	}
	prOpt := metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(labelMap).String(),
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
