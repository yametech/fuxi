package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"k8s.io/api/core/v1"
	"net/http"
)

// Get PersistentVolumeClaims
func (w *WorkloadsAPI) GetPersistentVolumeClaims(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.persistentVolumeClaims.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// Delete PersistentVolumeClaims
func (w *WorkloadsAPI) DeletePersistentVolumeClaims(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	err := w.persistentVolumeClaims.Delete(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, "")
}

// List PersistentVolumeClaims
func (w *WorkloadsAPI) ListPersistentVolumeClaims(g *gin.Context) {
	list, err := resourceList(g, w.persistentVolumeClaims)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	persistentVolumeClaimsList := &v1.PersistentVolumeClaimList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, persistentVolumeClaimsList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, persistentVolumeClaimsList)
}
