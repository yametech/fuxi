package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ListCustomResourceDefinition List CustomResourceDefinition
func (w *WorkloadsAPI) ListCustomResourceDefinition(g *gin.Context) {
	list, _ := w.customResourceDefinition.List(dyn.ResourceCustomResourceDefinition, "", "", 0, 0, nil)
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

// ListGeneralCustomResourceDefinition List General CustomResourceDefinition
func (w *WorkloadsAPI) ListGeneralCustomResourceDefinition(g *gin.Context) {
	// crd param
	group := g.Param("group")
	version := g.Param("version")
	resource := g.Param("resource")

	//  import general GroupVersionResource
	groupVersionResource := schema.GroupVersionResource{Group: group, Version: version, Resource: resource}
	list, _ := w.customResourceDefinition.List(groupVersionResource, "", "", 0, 0, nil)
	customResourceDefinitionList := &v1beta1.CustomResourceDefinitionList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	_ = json.Unmarshal(marshalData, customResourceDefinitionList)
	g.JSON(http.StatusOK, list)
}
