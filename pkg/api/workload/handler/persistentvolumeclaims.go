package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"k8s.io/api/core/v1"
	"net/http"
)

// Get PersistentVolumeClaims
func (w *WorkloadsAPI) GetPersistentVolumeClaims(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.persistentVolumeClaims.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List PersistentVolumeClaims
func (w *WorkloadsAPI) ListPersistentVolumeClaims(g *gin.Context) {
	list, err := w.persistentVolumeClaims.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	persistentVolumeClaimsList := &v1.PersistentVolumeClaimList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, persistentVolumeClaimsList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, persistentVolumeClaimsList)
}
