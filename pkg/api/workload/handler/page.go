package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

// Get Page
func (w *WorkloadsAPI) GetPage(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.page.Get(namespace, name)
	if err != nil {
		toRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Page
func (w *WorkloadsAPI) ListPage(g *gin.Context) {
	list, err := w.form.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	pageList := &v1.PageList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, pageList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, pageList)
}

// Create Page
func (w *WorkloadsAPI) CreatePage(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		toRequestParamsError(g, err)
		return
	}

	obj := v1.Form{}
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
	newObj, err := w.page.Apply(obj.Namespace, obj.Name, unstructuredStruct)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}
