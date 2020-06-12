package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/pkg/api/common"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"

	v1 "github.com/alauda/kube-ovn/pkg/apis/kubeovn/v1"
	"github.com/gin-gonic/gin"
)

// Get SubNet
func (w *WorkloadsAPI) GetSubNet(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.subnet.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List SubNet
func (w *WorkloadsAPI) ListSubNet(g *gin.Context) {
	list, err := resourceList(g, w.ip)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	roleList := &v1.SubnetList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, roleList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, roleList)
}

// Create SubNet
func (w *WorkloadsAPI) CreateSubNet(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := v1.Subnet{}
	err = json.Unmarshal(rawData, &obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	newObj, err := w.subnet.Apply(obj.Namespace, obj.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}
