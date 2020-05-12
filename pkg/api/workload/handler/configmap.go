package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"k8s.io/api/core/v1"
	"net/http"
)

// Get ConfigMaps
func (w *WorkloadsAPI) GetConfigMaps(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.configMaps.Get(namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ConfigMaps
func (w *WorkloadsAPI) ListConfigMaps(g *gin.Context) {
	list, _ := w.configMaps.List("", "", 0, 100, nil)
	configMapList := &v1.ConfigMapList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, configMapList)
	g.JSON(http.StatusOK, configMapList)
}
