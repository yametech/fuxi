package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	nuwav1 "github.com/yametech/nuwa/api/v1"
	"net/http"
)

// Get Stone
func (w *WorkloadsAPI) GetStone(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")

	item, err := w.stone.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List StatefulSet
func (w *WorkloadsAPI) ListStone(g *gin.Context) {
	list, err := resourceList(g, w.stone)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	stoneList := &nuwav1.StoneList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, stoneList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, stoneList)
}
