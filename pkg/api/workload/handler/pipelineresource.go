package handler

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"

	"github.com/gin-gonic/gin"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/yametech/fuxi/pkg/api/common"
)

func (w *WorkloadsAPI) GetPipelineResource(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.pipelineResource.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

func (w *WorkloadsAPI) ListPipelineResource(g *gin.Context) {
	var list *unstructured.UnstructuredList
	var err error
	namespace := g.Param("namespace")
	if namespace == "" {
		list, err = w.pipelineResource.List("", "", 0, 0, nil)
	} else {
		labelSelector := fmt.Sprintf("namespace=%s", namespace)
		list, err = w.pipelineResource.List("", "", 0, 0, labelSelector)
	}
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	pipelineResourceList := &tekton.PipelineResourceList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, pipelineResourceList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, pipelineResourceList)
}

// Create PipelineResource
func (w *WorkloadsAPI) CreatePipelineResource(g *gin.Context) {
	namespace := g.Param("namespace")
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := tekton.PipelineResource{}
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
	newObj, err := w.pipelineResource.Apply(namespace, obj.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}
