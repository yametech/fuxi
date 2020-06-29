package donkey

import (
	"encoding/json"
	"fmt"
	fuxi "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
	"github.com/yametech/fuxi/pkg/service/workload"
	v1 "k8s.io/api/core/v1"
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

// New Secret Object
func NewSecretObj(namespace string, name string, register fuxi.Stack) (*unstructured.Unstructured, error) {
	obj := &v1.Secret{}
	obj.Namespace = namespace

	if register.Verification == "Account" {
		item := map[string]map[string]string{register.Address: {"username": register.User, "password": register.Password}}
		itemString, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		obj.Type = "kubernetes.io/dockercfg"
		obj.Data = map[string][]byte{".dockercfg": itemString}
		obj.Name = fmt.Sprint(name, "-dockercfg", "-secret")
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		return nil, err
	}
	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	return workload.NewSecrets().Apply(obj.Namespace, obj.Name, unstructuredStruct)
}

// New Service Accounts Object
func NewServiceAccountsObj(namespace string, secretName string) (*unstructured.Unstructured, error) {
	obj := &v1.ServiceAccount{}
	obj.Namespace = namespace
	obj.Name = "fuxi"
	obj.Secrets = append(obj.Secrets, v1.ObjectReference{Name: secretName, Namespace: namespace})

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		return nil, err
	}
	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	return workload.NewServiceAccount().Apply(obj.Namespace, obj.Name, unstructuredStruct)
}

// Bulk Create Service Accounts
func BulkCreateServiceAccounts(dept fuxi.BaseDepartment) error {
	if dept.Spec.Registers != nil {
		for rIndex := range dept.Spec.Registers {
			for nIndex := range dept.Spec.Namespace {
				newSecret, err := NewSecretObj(
					dept.Spec.Namespace[nIndex], dept.Name, dept.Spec.Registers[rIndex])
				if err != nil {
					return err
				}
				_, err = NewServiceAccountsObj(newSecret.GetNamespace(), newSecret.GetName())
				if err != nil {
					return err
				}
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
				if err := BulkCreateServiceAccounts(*dept); err != nil {
					return fmt.Errorf(err.Error())
				}
				d.update(dept)
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

func (d *DepartmentAssistant) update(dept *fuxi.BaseDepartment) {

}

func runtimeObjectToInstanceObj(robj runtime.Object, targeObj interface{}) error {
	bytesData, err := json.Marshal(robj)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytesData, targeObj)
}
