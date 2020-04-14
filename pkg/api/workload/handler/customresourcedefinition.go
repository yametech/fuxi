package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"net/http"
)

// Get CustomResourceDefinition
func (w *WorkloadsAPI) GetCustomResourceDefinition(g *gin.Context) {
	name := g.Param("name")
	item, err := w.customResourceDefinition.Get(dyn.ResourceCustomResourceDefinition, "", name)
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
