package handler

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"net/http"

	"github.com/gin-gonic/gin"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/yametech/fuxi/pkg/api/common"
)

func (w *WorkloadsAPI) GetTaskRun(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.taskRun.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

func (w *WorkloadsAPI) ListTaskRun(g *gin.Context) {
	var list *unstructured.UnstructuredList
	var err error
	namespace := g.Param("namespace")
	if namespace == "" {
		list, err = w.taskRun.List("", "", 0, 0, nil)
	} else {
		labelSelector := fmt.Sprintf("namespace=%s", namespace)
		list, err = w.taskRun.List("", "", 0, 0, labelSelector)
	}
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	taskRunList := &tekton.TaskRunList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, taskRunList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, taskRunList)
}
