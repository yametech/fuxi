package ops

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
)

type ResourceService interface {
	CreateOrUpdatePipelineResource(resource Resource) error
	PipelineResourceDelete(namespace, name string) error
	PipelineResourceList(namespace string) ([]Resource, error)
	GetPipelineResource(namespace, name string) (*Resource, error)
}

type Resource struct {
	Name         string
	Namespace    string
	Lables       map[string]string
	ResourceType v1alpha1.PipelineResourceType
	Params       []v1alpha1.ResourceParam
	SecretParams []v1alpha1.SecretParam
}

//
//func toPipelineResource(rs *Resource) *v1alpha1.PipelineResource {
//	return &v1alpha1.PipelineResource{
//		TypeMeta: metav1.TypeMeta{},
//		ObjectMeta: metav1.ObjectMeta{
//			Name:   rs.Name,
//			Labels: rs.Lables,
//		},
//		Spec: v1alpha1.PipelineResourceSpec{
//			Type:         rs.ResourceType,
//			Params:       rs.Params,
//			SecretParams: nil,
//		},
//		//Status: v1alpha1.PipelineResourceStatus{},
//	}
//}
//
//func (ops *Ops) getPipelineResources(name, namespace string) (*v1alpha1.PipelineResource, error) {
//	if name == "" && namespace == "" {
//		return nil, errors.New("name or namespace should not be empty")
//	}
//	res, err := ops.client.PipelineResources(namespace).
//		Get(name, metav1.GetOptions{})
//	if err != nil {
//		return res, err
//	}
//	return res, err
//}
//
////CreateOrUpdateResource creates or updates PipelineResource
//func (ops *Ops) CreateOrUpdatePipelineResource(resource Resource) error {
//
//	oldGitPipelineResource, err := ops.getPipelineResources(resource.Name, resource.Namespace)
//	if err != nil {
//		if kubeeror.IsNotFound(err) {
//			_, err := ops.client.
//				PipelineResources(resource.Namespace).
//				Create(toPipelineResource(&resource))
//			if err != nil {
//				return err
//			}
//		} else {
//			return err
//		}
//
//	}
//
//	if oldGitPipelineResource.Spec.DeepCopy().Type != "" {
//		newGitPipelineResource := toPipelineResource(&resource)
//		if !reflect.DeepEqual(oldGitPipelineResource.Spec, newGitPipelineResource.Spec) {
//			oldGitPipelineResource.Spec = newGitPipelineResource.Spec
//			_, err = ops.client.PipelineResources(resource.Namespace).Update(oldGitPipelineResource)
//			if err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
//
////PipelineResourceDelete delete a pipelineresource  resource
//func (ops *Ops) PipelineResourceDelete(namespace, name string) error {
//	if name == "" && namespace == "" {
//		return errors.New("name or namespace should not be empty")
//	}
//	err := ops.client.PipelineResources(namespace).Delete(name, &metav1.DeleteOptions{})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////PipelineResourceList list current namespace/department all data
//func (ops *Ops) PipelineResourceList(namespace string) ([]Resource, error) {
//
//	if namespace == "" {
//		return nil, errors.New("namespace should not be empty")
//	}
//
//	var strVals []string
//	strVals = append(strVals, namespace)
//	key := "namespace"
//	rq, err := labels.NewRequirement(key, selection.Equals, strVals)
//	if err != nil {
//		return nil, err
//	}
//
//	lable := labels.NewSelector().Add(*rq)
//	rs, err := ops.informer.Tekton().V1alpha1().PipelineResources().
//		Lister().
//		PipelineResources(namespace).List(lable)
//
//	if err != nil {
//		return nil, err
//	}
//
//	var res []Resource
//	for _, ps := range rs {
//		res = append(res, Resource{
//			Name:         ps.Name,
//			Lables:       nil,
//			ResourceType: ps.Spec.Type,
//			Params:       ps.Spec.Params,
//			SecretParams: ps.Spec.SecretParams,
//		})
//	}
//
//	return res, nil
//}
//
////PipelineResourceDelete delete a pipeline resource
//func (ops *Ops) GetPipelineResource(namespace, name string) (*Resource, error) {
//	res, err := ops.getPipelineResources(name, namespace)
//	if err != nil {
//		return nil, err
//	}
//	return &Resource{
//		Name:         res.Name,
//		Lables:       res.Labels,
//		ResourceType: res.Spec.Type,
//		Params:       res.Spec.Params,
//		SecretParams: res.Spec.SecretParams,
//	}, nil
//}
