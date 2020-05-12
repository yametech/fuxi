package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/rbac/v1"
	"net/http"
)

// Get RoleBinding
func (w *WorkloadsAPI) GetRoleBinding(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.roleBinding.Get(namespace, name)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List RoleBinding
func (w *WorkloadsAPI) ListRoleBinding(g *gin.Context) {
	list, err := w.roleBinding.List("", "", 0, 10000, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	roleBindingList := &v1.RoleBindingList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, roleBindingList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, roleBindingList)
}
