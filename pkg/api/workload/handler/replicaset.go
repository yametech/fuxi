package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	"net/http"
)

// Get ReplicaSet
func (w *WorkloadsAPI) GetReplicaSet(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.replicaSet.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ReplicaSet
func (w *WorkloadsAPI) ListReplicaSet(g *gin.Context) {
	list, err := w.replicaSet.List("", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	replicaSetList := &v1.ReplicaSetList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, replicaSetList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, replicaSetList)
}
