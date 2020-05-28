package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"net/http"
)

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
