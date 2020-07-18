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

// Get TektonGraph
func (w *WorkloadsAPI) GetTektonGraph(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.tektonGraph.Get(namespace, name)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List TektonGraph
func (w *WorkloadsAPI) ListTektonGraph(g *gin.Context) {
	list, err := resourceList(g, w.tektonGraph)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	tektonGraphList := &v1.TektonGraphList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, tektonGraphList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, tektonGraphList)
}

// Create TektonGraph
func (w *WorkloadsAPI) CreateTektonGraph(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := v1.TektonGraph{}
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
	newObj, _, err := w.tektonGraph.Apply(obj.Namespace, obj.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}
