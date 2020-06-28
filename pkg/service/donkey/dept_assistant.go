package donkey

import (
	"encoding/json"
	"fmt"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
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

func (d *DepartmentAssistant) Run() error {
	d.SetGroupVersionResource(types.ResourceBaseDepartment)
	list, err := d.List("", "", 0, 0, nil)
	if err != nil {
		return err
	}
	baseDepartmentList := &v1.BaseDepartmentList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshalData, baseDepartmentList)
	if err != nil {
		return err
	}
	for _, item := range baseDepartmentList.Items {
		d.update(item)
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
				dept := &v1.BaseDepartment{}
				runtimeObjectToInstanceObj(item.Object, dept)
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

func (d *DepartmentAssistant) update(dept *v1.BaseDepartment) {

}

func runtimeObjectToInstanceObj(robj runtime.Object, targeObj interface{}) error {
	bytesData, err := json.Marshal(robj)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytesData, targeObj)
}
