package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Secret
func (w *WorkloadsAPI) GetSecret(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.secret.Get(dyn.ResourceSecrets, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Secret
func (w *WorkloadsAPI) ListSecret(g *gin.Context) {
	list, _ := w.secret.List(dyn.ResourceSecrets, "", "", 0, 10000, nil)
	secretList := &v1.SecretList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, secretList)
	g.JSON(http.StatusOK, secretList)
}
