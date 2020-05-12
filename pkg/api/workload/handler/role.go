package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/rbac/v1"
	"net/http"
)

// Get Role
func (w *WorkloadsAPI) GetRole(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.role.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Role
func (w *WorkloadsAPI) ListRole(g *gin.Context) {
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
