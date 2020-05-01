package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

// Create Namespace
func (w *WorkloadsAPI) CreateNamespace(g *gin.Context) {
	version := g.Param("version")
	if version != "v1" {
		g.JSON(
			http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    "url version is not v1",
				status: "Request bad parameter",
			})
		return
	}
	rawData, err := g.GetRawData()
	if version != "v1" {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   string(rawData),
				msg:    "post data error",
				status: "Request bad parameter",
			})
		return
	}
	obj := corev1.Namespace{}
	err = json.Unmarshal(rawData, &obj)
	if err != nil {
		g.JSON(
			http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   string(rawData),
				msg:    "unmarshal post data error",
				status: "Request bad parameter",
			})
		return
	}
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		g.JSON(
			http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter",
			})
		return
	}
	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	newObj, err := w.namespace.Apply(
		dyn.ResourceNamespace,
		"",
		obj.Name,
		unstructuredStruct,
	)
	if err != nil {
		g.JSON(
			http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   newObj,
				msg:    err.Error(),
				status: "apply namespace error",
			})
		return
	}

	g.JSON(http.StatusOK, newObj)
}

// Delete Namespace
func (w *WorkloadsAPI) DeleteNamespace(g *gin.Context) {
	version := g.Param("version")
	if version != "v1" {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    "url version is not v1",
				status: "Request bad parameter",
			})
		return
	}
	namespaceName := g.Param("namespace")
	err := w.namespace.Delete(dyn.ResourceNamespace, "", namespaceName)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter",
			})
		return
	}
	g.JSON(http.StatusOK, "")
}

// Get Namespace
func (w *WorkloadsAPI) GetNamespace(g *gin.Context) {
	namespaceName := g.Param("namespace")
	item, err := w.namespace.Get(dyn.ResourceNamespace, "", namespaceName)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Namespaces
func (w *WorkloadsAPI) ListNamespace(g *gin.Context) {
	list, _ := w.namespace.List(dyn.ResourceNamespace, "", "", 0, 10000, nil)
	namespaceList := &corev1.NamespaceList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, namespaceList)
	g.JSON(http.StatusOK, namespaceList)
}
