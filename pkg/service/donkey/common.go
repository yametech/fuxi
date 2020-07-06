package donkey

import (
	"encoding/json"
	constraint "github.com/yametech/fuxi/common"
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// Get: secretList
func SecretList() (*v1.SecretList, error) {
	d := NewDepartmentAssistant()
	d.SetGroupVersionResource(types.ResourceSecrets)
	list, _ := d.List("", "", 0, 0, nil)

	secretList := &v1.SecretList{}
	marshalData, _ := json.Marshal(list)
	if err := json.Unmarshal(marshalData, secretList); err != nil {
		return nil, err
	}
	return secretList, nil
}

// Get: bindingAnnotations
func GetBindingAnnotations(secret *v1.Secret) (string, bool) {
	annotations := secret.GetAnnotations()
	if annotations == nil {
		return "", false
	}

	binding, ok := annotations[constraint.DepartmentBindingSecret]
	if !ok {
		return "", false
	}
	return binding, true
}

// Get: secret department array annotations
func GetDepartmentAnnotations(secret *v1.Secret) []string {
	var secretAnnotations = make([]string, 0)
	binding, ok := GetBindingAnnotations(secret)
	if !ok {
		return secretAnnotations
	}

	if err := json.Unmarshal([]byte(binding), &secretAnnotations); err == nil {
		return secretAnnotations
	}
	return secretAnnotations
}

// Func: unique department annotations
func UniqueDepartmentAnnotations(annotations []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for _, entry := range annotations {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// Func: LocalObjectReference deduplication
func UniqueImagePullSecrets(mapSlice []v1.LocalObjectReference) []v1.LocalObjectReference {
	keys := make(map[string]bool)
	list := make([]v1.LocalObjectReference, 0)
	for _, entry := range mapSlice {
		if _, value := keys[entry.Name]; !value {
			keys[entry.Name] = true
			list = append(list, entry)
		}
	}
	return list
}

// Func: set secret annotations
func SetSecretAnnotations(secret *v1.Secret, departments []string) {

	marshalData, _ := json.Marshal(departments)
	annotations := secret.GetAnnotations()
	if annotations == nil {
		secret.ObjectMeta.Annotations = map[string]string{
			constraint.DepartmentBindingSecret: string(marshalData),
		}
	} else {
		annotations[constraint.DepartmentBindingSecret] = string(marshalData)
		secret.ObjectMeta.Annotations = annotations
	}

	unstructuredObj, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(&secret)
	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}

	b := NewDepartmentAssistant()
	_, _ = b.Apply(secret.GetNamespace(), secret.GetName(), unstructuredStruct)
}

// set department array secretAnnotations
func SetSecretDepartmentAnnotations(secret *v1.Secret, name string) {
	secretDepartmentAnnotations := GetDepartmentAnnotations(secret)

	list := UniqueDepartmentAnnotations(secretDepartmentAnnotations)
	exists := false
	for i := range list {
		if list[i] == name {
			exists = true
		}
	}

	if !exists {
		secretDepartmentAnnotations = append(secretDepartmentAnnotations, name)
		SetSecretAnnotations(secret, secretDepartmentAnnotations)
	}
}

// Func: find department array secret annotations
func FindSecretAnnotationsByDepartment(department string) ([]v1.Secret, error) {
	secretList, err := SecretList()
	if err != nil {
		return nil, err
	}

	var summoner []v1.Secret
	for _, item := range secretList.Items {
		binding := make([]string, 0)
		bindingString, ok := GetBindingAnnotations(&item)
		if ok {
			err := json.Unmarshal([]byte(bindingString), &binding)
			if err == nil {
				for i := range binding {
					if binding[i] == department {
						summoner = append(summoner, item)
					}
				}
			}
		}
	}
	return summoner, nil
}

// Func: remove secretAnnotations by department
func RemoveSecretAnnotationsDepartment(secret v1.Secret, departmentNamespace []string) {

	for i := range departmentNamespace {
		if departmentNamespace[i] == secret.GetNamespace() {
			departmentNamespace = append(departmentNamespace[:i], departmentNamespace[i+1:]...)
		}
	}

	if len(departmentNamespace) > 0 {
		annotations := secret.GetAnnotations()
		if annotations != nil {
			binding := make([]string, 0)
			bindingString, ok := annotations[constraint.DepartmentBindingSecret]
			if ok {
				err := json.Unmarshal([]byte(bindingString), &binding)
				if err == nil {

				}
			}
		}
	}

}
