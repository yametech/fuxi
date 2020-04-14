package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/api/core/v1"
	"net/http"
)

// Get PersistentVolumeClaims
func (w *WorkloadsAPI) GetPersistentVolumeClaims(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.persistentVolumeClaims.Get(dyn.ResourcePersistentVolumeClaims, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List PersistentVolumeClaims
func (w *WorkloadsAPI) ListPersistentVolumeClaims(g *gin.Context) {
	list, _ := w.persistentVolumeClaims.List(dyn.ResourcePersistentVolumeClaims, "", "", 0, 100, nil)
	persistentVolumeClaimsList := &v1.PersistentVolumeClaimList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, persistentVolumeClaimsList)
	g.JSON(http.StatusOK, persistentVolumeClaimsList)
}
