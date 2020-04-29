package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "github.com/yametech/nuwa/api/v1"
	"net/http"
)

// Get StatefulSet
func (w *WorkloadsAPI) GetStatefulSet(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")

	item, err := w.statefulSet.Get(dyn.ResourceStatefulSet, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List StatefulSet
func (w *WorkloadsAPI) ListStatefulSet(g *gin.Context) {
	list, _ := w.statefulSet.List(dyn.ResourceStatefulSet, "", "", 0, 0, nil)
	statefulSetList := &v1.StatefulSetList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, statefulSetList)
	g.JSON(http.StatusOK, statefulSetList)
}
