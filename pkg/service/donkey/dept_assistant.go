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
		if err := common.RuntimeObjectToInstanceObj(secret, obj); err != nil {
			return nil, err
		}
	}

	// Annotation department information
	SetDepartmentAnnotations(obj, name)
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
	obj := &v1.ServiceAccount{}
	d.SetGroupVersionResource(types.ResourceServiceAccount)
	serviceAccount, err := d.Get(namespace, "default")

	if errors.IsNotFound(err) {
		obj = &v1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "default",
				Namespace: namespace,
			},
		}
	} else {
		if err := common.RuntimeObjectToInstanceObj(serviceAccount, obj); err != nil {
			return err
		}
	}

	ImagePullSecrets := UniqueImagePullSecrets(obj.ImagePullSecrets)

	secretObjects := make([]map[string]string, 0)
	for _, alreadySecret := range ImagePullSecrets {
		secretObjects = append(secretObjects, map[string]string{"name": alreadySecret.Name})
	}

	// If secretName does not exist, add it to secretObjects
	contains := false
	for i := range secretObjects {
		if secretObjects[i]["name"] == secretName {
			contains = true
		}
	}

	if contains == false {
		secretObjects = append(secretObjects, map[string]string{"name": secretName})
	}

	patchData := map[string]interface{}{
		"imagePullSecrets": secretObjects,
	}

	d.SetGroupVersionResource(types.ResourceServiceAccount)
	_, err = d.Patch(namespace, "default", patchData)
	return nil
}

// bulkCreateServiceAccounts Create Service Accounts
func (d *DepartmentAssistant) bulkPatchServiceAccounts(dept *fuxi.BaseDepartment) error {
	if dept.Spec.Registers == nil {
		return nil
	}

	// sync secret department annotations
	if summoner, err := FindSecretAnnotationsByDepartment(dept.Name); err == nil {
		fmt.Print("summoner", summoner, "\n")
		fmt.Print("Namespace", dept.Spec.Namespace, "\n")
		for i := range summoner {
			SyncDepartmentAnnotations(&summoner[i], dept.Spec.Namespace)
		}
	}

	// clear service accounts with empty register

	// clear namespace deleted service account
	for rIndex := range dept.Spec.Registers {
		for nIndex := range dept.Spec.Namespace {
			fmt.Print("namespace ", dept.Spec.Namespace[nIndex], "\n")

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
				if err := common.RuntimeObjectToInstanceObj(item.Object, dept); err != nil {
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
