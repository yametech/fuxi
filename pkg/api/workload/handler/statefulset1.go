package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	nuwav1 "github.com/yametech/nuwa/api/v1"
	"net/http"
)

// Get StatefulSet
func (w *WorkloadsAPI) GetStatefulSet1(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")

	item, err := w.statefulSet1.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List StatefulSet
func (w *WorkloadsAPI) ListStatefulSet1(g *gin.Context) {
	list, err := w.statefulSet1.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	statefulSetList := &nuwav1.StatefulSetList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, statefulSetList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, statefulSetList)
}
