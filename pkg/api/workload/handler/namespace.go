package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

// Create Namespace
func (w *WorkloadsAPI) CreateNamespace(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := corev1.Namespace{}
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
	newObj, err := w.namespace.Apply("", obj.Name, unstructuredStruct)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}

// Delete Namespace
func (w *WorkloadsAPI) DeleteNamespace(g *gin.Context) {
	namespaceName := g.Param("namespace")
	err := w.namespace.Delete("", namespaceName)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, "")
}

// Get Namespace
func (w *WorkloadsAPI) GetNamespace(g *gin.Context) {
	namespaceName := g.Param("namespace")
	item, err := w.namespace.Get("", namespaceName)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	g.JSON(http.StatusOK, item)
}

// List Namespaces
func (w *WorkloadsAPI) ListNamespace(g *gin.Context) {
	list, err := w.namespace.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	namespaceList := &corev1.NamespaceList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	err = json.Unmarshal(marshalData, namespaceList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, namespaceList)
}
