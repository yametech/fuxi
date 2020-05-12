package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

// Get Secret
func (w *WorkloadsAPI) GetSecret(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.secret.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Secret
func (w *WorkloadsAPI) ListSecret(g *gin.Context) {
	list, err := w.secret.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	secretList := &v1.SecretList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, secretList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, secretList)
}
