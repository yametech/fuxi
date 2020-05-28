package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"net/http"
)

// Get FormRender
func (w *WorkloadsAPI) GetFormRender(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.formRender.Get(namespace, name)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List FormRender
func (w *WorkloadsAPI) ListFormRender(g *gin.Context) {
	list, err := w.formRender.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	formRenderList := &v1.FormRenderList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, formRenderList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, formRenderList)
}
