package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
)

// Get CustomResourceDefinition
func (w *WorkloadsAPI) GetCustomResourceDefinition(g *gin.Context) {
	// crd param
	group := g.Param("group")
	version := g.Param("version")
	resource := g.Param("resource")
	//  import general GroupVersionResource
	groupVersionResource := schema.GroupVersionResource{Group: group, Version: version, Resource: resource}
	namespace := g.Query("namespace")
	name := g.Query("name")
	item, err := w.customResourceDefinition.Get(groupVersionResource, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List CustomResourceDefinition
func (w *WorkloadsAPI) ListCustomResourceDefinition(g *gin.Context) {
	list, _ := w.customResourceDefinition.List(dyn.ResourceCustomResourceDefinition, "", "", 0, 10000, nil)
	customResourceDefinitionList := &v1beta1.CustomResourceDefinitionList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, customResourceDefinitionList)
	g.JSON(http.StatusOK, customResourceDefinitionList)
}

// List General CustomResourceDefinition
func (w *WorkloadsAPI) ListGeneralCustomResourceDefinition(g *gin.Context) {
	// crd param
	group := g.Param("group")
	version := g.Param("version")
	resource := g.Param("resource")
	//  import general GroupVersionResource
	groupVersionResource := schema.GroupVersionResource{Group: group, Version: version, Resource: resource}
	list, _ := w.customResourceDefinition.List(groupVersionResource, "", "", 0, 10000, nil)
	customResourceDefinitionList := &v1beta1.CustomResourceDefinitionList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, customResourceDefinitionList)
	g.JSON(http.StatusOK, list)
}
