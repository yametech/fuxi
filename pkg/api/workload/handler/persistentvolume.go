package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"k8s.io/api/core/v1"
	"net/http"
)

// Get PersistentVolume
func (w *WorkloadsAPI) GetPersistentVolume(g *gin.Context) {
	name := g.Param("name")
	item, err := w.persistentVolume.Get("", name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// Delete PersistentVolume
func (w *WorkloadsAPI) DeletePersistentVolume(g *gin.Context) {
	name := g.Param("name")
	err := w.persistentVolume.Delete("", name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, "")
}

// List PersistentVolume
func (w *WorkloadsAPI) ListPersistentVolume(g *gin.Context) {
	list, err := resourceList(g, w.persistentVolume)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	persistentVolumeList := &v1.PersistentVolumeList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, persistentVolumeList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, persistentVolumeList)
}
