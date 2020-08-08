package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

// Get TektonStore
func (w *WorkloadsAPI) GetTektonStore(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.tektonStore.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List TektonStore
func (w *WorkloadsAPI) ListTektonStore(g *gin.Context) {
	list, err := resourceList(g, w.tektonStore)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	TektonStoreList := &v1.TektonStoreList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, TektonStoreList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, TektonStoreList)
}

// Create TektonStore
func (w *WorkloadsAPI) CreateTektonStore(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := v1.TektonStore{}
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
	newObj, _, err := w.tektonStore.Apply(obj.Namespace, obj.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}
