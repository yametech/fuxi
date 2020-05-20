package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/rbac/v1"
)

// Get ClusterRole
func (w *WorkloadsAPI) GetClusterRole(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.role.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ClusterRole
func (w *WorkloadsAPI) ListClusterRole(g *gin.Context) {
	list, err := w.role.List("", "", 0, 10000, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	roleList := &v1.RoleList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, roleList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, roleList)
}
