package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"k8s.io/api/storage/v1"
	"net/http"
)

// Get StorageClass
func (w *WorkloadsAPI) GetStorageClass(g *gin.Context) {
	name := g.Param("name")
	item, err := w.storageClass.Get("", name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List StorageClass
func (w *WorkloadsAPI) ListStorageClass(g *gin.Context) {
	list, err := w.storageClass.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	storageClassList := &v1.StorageClassList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, storageClassList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, storageClassList)
}
