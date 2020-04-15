package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/api/storage/v1"
	"net/http"
)

// Get StorageClass
func (w *WorkloadsAPI) GetStorageClass(g *gin.Context) {
	name := g.Param("name")
	item, err := w.storageClass.Get(dyn.ResourceStorageClass, "", name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List StorageClass
func (w *WorkloadsAPI) ListStorageClass(g *gin.Context) {
	list, _ := w.storageClass.List(dyn.ResourceStorageClass, "", "", 0, 0, nil)
	storageClassList := &v1.StorageClassList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, storageClassList)
	g.JSON(http.StatusOK, storageClassList)
}
