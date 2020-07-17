package handler

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"

	"github.com/gin-gonic/gin"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/yametech/fuxi/pkg/api/common"
	service_common "github.com/yametech/fuxi/pkg/service/common"
)

func (w *WorkloadsAPI) UpdatePipeline(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	tempPipeline := &tekton.Pipeline{}
	if err := json.Unmarshal(rawData, tempPipeline); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	pipeline, err := w.pipeline.Get(namespace, name)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	pipelineObject := &tekton.Pipeline{}
	if err := service_common.RuntimeObjectToInstanceObj(pipeline, pipelineObject); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	pipelineObject.ObjectMeta = tempPipeline.ObjectMeta
	pipelineObject.Spec = tempPipeline.Spec

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&pipelineObject)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	newObj, err := w.pipeline.Apply(namespace, name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}

func (w *WorkloadsAPI) CreatePipeline(g *gin.Context) {
	namespace := g.Param("namespace")
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := tekton.Pipeline{}
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
	newObj, err := w.pipeline.Apply(namespace, obj.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}

// Get Pipeline
func (w *WorkloadsAPI) GetPipeline(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.pipeline.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Pipeline
func (w *WorkloadsAPI) ListPipeline(g *gin.Context) {
	list, err := resourceList(g, w.pipeline)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	pipelineList := &tekton.PipelineList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, pipelineList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, pipelineList)
}
