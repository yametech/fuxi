package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/rbac/v1"
)

// Get Role
func (w *WorkloadsAPI) GetRole(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.role.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ClusterRole
func (w *WorkloadsAPI) ListRole(g *gin.Context) {
	list, err := resourceList(g, w.role)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	roleList := &v1.RoleList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, roleList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, roleList)
}
