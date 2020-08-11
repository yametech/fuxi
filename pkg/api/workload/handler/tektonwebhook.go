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

// Get TektonWebHook
func (w *WorkloadsAPI) GetTektonWebHook(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.tektonWebHook.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List TektonWebHook
func (w *WorkloadsAPI) ListTektonWebHook(g *gin.Context) {
	list, err := resourceList(g, w.tektonWebHook)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	tektonWebHookList := &v1.TektonWebHookList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, tektonWebHookList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, tektonWebHookList)
}

// Create TektonWebHook
func (w *WorkloadsAPI) CreateTektonWebHook(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := v1.TektonWebHook{}
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
	newObj, _, err := w.tektonWebHook.Apply(obj.Namespace, obj.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}
