package donkey

import (
	"encoding/json"
	"fmt"
	fuxi "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

type DepartmentAssistant struct {
	common.WorkloadsResourceHandler
	stopChan chan struct{}
}

// NewDepartmentAssistant exported
func NewDepartmentAssistant() *DepartmentAssistant {
	return &DepartmentAssistant{
		&common.DefaultImplWorkloadsResourceHandler{},
		make(chan struct{}),
	}
}

// new Secret Object
func (d *DepartmentAssistant) updateSecretObject(namespace string, name string, register *fuxi.Stack) (*unstructured.Unstructured, error) {
	obj := &v1.Secret{}
	newName := fmt.Sprintf("%s-%s-%s", name, "dockerconfigjson", "secret")
	d.SetGroupVersionResource(types.ResourceSecrets)
	secret, err := d.Get(namespace, newName)

	if errors.IsNotFound(err) {
		obj = &v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      newName,
				Namespace: namespace,
			},
			Type: "kubernetes.io/dockerconfigjson",
		}
	} else {
		if err := runtimeObjectToInstanceObj(secret, obj); err != nil {
			return nil, err
		}
	}

	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	if register.Verification == "Account" {
		// eg: {"auths":{"registry.cn-shenzhen.aliyuncs.com":{"username":"us","password":"pwd","email":"laik.lj@me.com","auth":"dsadsada"}}}
		bytesData, err := common.HandleDockerCfgJSONContent(register.User, register.Password, "yame@yame.com", register.Address)
		if err != nil {
			return nil, err
		}
		obj.Data = map[string][]byte{".dockerconfigjson": bytesData}

	}
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		return nil, err
	}
	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	d.SetGroupVersionResource(types.ResourceSecrets)
	return d.Apply(obj.Namespace, obj.Name, unstructuredStruct)
}

// patchServiceAccount
func (d *DepartmentAssistant) patchServiceAccount(namespace string, secretName string) error {
	d.SetGroupVersionResource(types.ResourceServiceAccount)
	serviceAccount, err := d.Get(namespace, "default")
	if err != nil {
		return err
	}
	obj := &v1.ServiceAccount{}
	if err := runtimeObjectToInstanceObj(serviceAccount, obj); err != nil {
		return err
	}
	obj.ImagePullSecrets = append(obj.ImagePullSecrets, v1.LocalObjectReference{Name: secretName})

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		return err
	}
	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	d.SetGroupVersionResource(types.ResourceServiceAccount)
	_, err = d.Apply(obj.Namespace, obj.Name, unstructuredStruct)
	if err != nil {
		return err
	}
	return nil
}

// bulkCreateServiceAccounts Create Service Accounts
func (d *DepartmentAssistant) bulkPatchServiceAccounts(dept *fuxi.BaseDepartment) error {
	if dept.Spec.Registers == nil {
		return nil
	}
	for rIndex := range dept.Spec.Registers {
		for nIndex := range dept.Spec.Namespace {
			newSecret, err := d.updateSecretObject(dept.Spec.Namespace[nIndex], dept.Name, &dept.Spec.Registers[rIndex])
			if err != nil {
				return err
			}
			err = d.patchServiceAccount(dept.Spec.Namespace[nIndex], newSecret.GetName())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *DepartmentAssistant) Run() error {
	d.SetGroupVersionResource(types.ResourceBaseDepartment)
	list, err := d.List("", "", 0, 0, nil)
	if err != nil {
		return err
	}
	baseDepartmentList := &fuxi.BaseDepartmentList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshalData, baseDepartmentList)
	if err != nil {
		return err
	}
	for _, item := range baseDepartmentList.Items {
		d.update(&item)
	}

	// watch
	d.SetGroupVersionResource(types.ResourceBaseDepartment)
	eventChan, err := d.Watch("", baseDepartmentList.ResourceVersion, 0, nil)
	if err != nil {
		return err
	}
	for {
		select {
		case item, ok := <-eventChan:
			if !ok {
				return fmt.Errorf("evnet watch error")
			}
			switch item.Type {
			case watch.Added, watch.Modified:
				dept := &fuxi.BaseDepartment{}
				if err := runtimeObjectToInstanceObj(item.Object, dept); err != nil {
					return err
				}
				if err := d.update(dept); err != nil {
					return err
				}
			}
		case <-d.stopChan:
			return nil
		}
	}
}

func (d *DepartmentAssistant) Stop() {
	d.stopChan <- struct{}{}
	return
}

func (d *DepartmentAssistant) update(dept *fuxi.BaseDepartment) error {
	if err := d.bulkPatchServiceAccounts(dept); err != nil {
		return err
	}
	return nil
}

func runtimeObjectToInstanceObj(robj runtime.Object, targeObj interface{}) error {
	bytesData, err := json.Marshal(robj)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytesData, targeObj)
}
