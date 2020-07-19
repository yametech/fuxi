package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	consts "github.com/yametech/fuxi/common"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func (w *WorkloadsAPI) PostWorkloadsTemplate(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	workloads := &v1.Workloads{}
	err = json.Unmarshal(rawData, workloads)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&workloads)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	newObj, _, err := w.workloadsTemplate.Apply("fuxi", workloads.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, newObj)
}

// Get WorkloadsTemplate
func (w *WorkloadsAPI) GetWorkloadsTemplate(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.workloadsTemplate.Get(namespace, name)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// list Shared WorkloadsTemplate
func (w *WorkloadsAPI) ListShareNamespacedWorkloadsTemplate(g *gin.Context) {
	namespace := g.Param("namespace")
	labelSelector := fmt.Sprintf("namespace=%s", namespace)

	list, err := w.workloadsTemplate.SharedNamespaceList(consts.WorkloadsDeployTemplateNamespace, labelSelector)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	workloadList := &v1.WorkloadsList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, workloadList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, workloadList)
}

// List WorkloadsTemplate
func (w *WorkloadsAPI) ListWorkloadsTemplate(g *gin.Context) {
	list, err := w.workloadsTemplate.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	workloadList := &v1.WorkloadsList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, workloadList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, workloadList)
}
