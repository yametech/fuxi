package common

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func UnstructuredObjectToInstanceObj(obj *unstructured.Unstructured, targetObj interface{}) error {
	bytesData, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytesData, targetObj)
}

func RuntimeObjectToInstanceObj(obj runtime.Object, targetObj interface{}) error {
	bytesData, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytesData, targetObj)
}

func InstanceToUnstructured(object runtime.Object) (*unstructured.Unstructured, error) {
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(object)
	if err != nil {
		return nil, err
	}

	return &unstructured.Unstructured{
		Object: unstructuredObj,
	}, nil
}
