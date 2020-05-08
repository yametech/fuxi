package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"net/http"
)

// Get FormRender
func (w *WorkloadsAPI) GetFormRender(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.formRender.Get(dyn.ResourceFormRender, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List FormRender
func (w *WorkloadsAPI) ListFormRender(g *gin.Context) {
	list, _ := w.formRender.List(dyn.ResourceFormRender, "", "", 0, 10000, nil)
	formRenderList := &v1.FormRenderList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, formRenderList)
	g.JSON(http.StatusOK, formRenderList)
}
