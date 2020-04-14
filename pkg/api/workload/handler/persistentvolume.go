package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/api/core/v1"
	"net/http"
)

// Get PersistentVolume
func (w *WorkloadsAPI) GetPersistentVolume(g *gin.Context) {
	name := g.Param("name")
	item, err := w.persistentVolume.Get(dyn.ResourcePersistentVolume, "", name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List PersistentVolume
func (w *WorkloadsAPI) ListPersistentVolume(g *gin.Context) {
	list, _ := w.persistentVolume.List(dyn.ResourcePersistentVolume, "", "", 0, 100, nil)
	persistentVolumeList := &v1.PersistentVolumeList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, persistentVolumeList)
	g.JSON(http.StatusOK, persistentVolumeList)
}
