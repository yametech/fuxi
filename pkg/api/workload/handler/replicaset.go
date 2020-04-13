package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/apps/v1"
	"net/http"
)

// Get ReplicaSet
func (w *WorkloadsAPI) GetReplicaSet(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.replicaSet.Get(dyn.ResourceReplicaSet, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ReplicaSet
func (w *WorkloadsAPI) ListReplicaSet(g *gin.Context) {
	list, _ := w.replicaSet.List(dyn.ResourceReplicaSet, "", "", 0, 10000, nil)
	replicaSetList := &v1.ReplicaSetList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, replicaSetList)
	g.JSON(http.StatusOK, replicaSetList)
}
