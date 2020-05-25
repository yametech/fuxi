package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

// Get Field
func (w *WorkloadsAPI) GetField(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.field.Get(namespace, name)
	if err != nil {
		toRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Field
func (w *WorkloadsAPI) ListField(g *gin.Context) {
	list, err := w.field.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	fieldList := &v1.FieldList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, fieldList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, fieldList)
}

// Create Field
func (w *WorkloadsAPI) CreateField(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		toRequestParamsError(g, err)
		return
	}

	obj := v1.Field{}
	err = json.Unmarshal(rawData, &obj)
	if err != nil {
		toRequestParamsError(g, err)
		return
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		toRequestParamsError(g, err)
		return
	}

	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	newObj, err := w.field.Apply(obj.Namespace, obj.Name, unstructuredStruct)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}
