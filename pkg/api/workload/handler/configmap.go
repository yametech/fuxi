package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/pkg/api/common"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
)

// Create ConfigMaps
func (w *WorkloadsAPI) CreateConfigMaps(g *gin.Context) {
	namespace := g.Param("namespace")
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := v1.ConfigMap{}
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

	newObj, err := w.configMaps.Apply(namespace, obj.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, newObj)
}

// Get ConfigMaps
func (w *WorkloadsAPI) GetConfigMaps(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.configMaps.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ConfigMaps
func (w *WorkloadsAPI) ListConfigMaps(g *gin.Context) {
	list, err := resourceList(g, w.configMaps)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	configMapList := &v1.ConfigMapList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, configMapList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, configMapList)
}
