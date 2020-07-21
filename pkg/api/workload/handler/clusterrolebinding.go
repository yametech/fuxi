package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"

	"github.com/gin-gonic/gin"
	rbacv1 "k8s.io/api/rbac/v1"
)

// Get ClusterRoleBind
func (w *WorkloadsAPI) GetClusterRoleBind(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.clusterRoleBinding.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ClusterRoleBind
func (w *WorkloadsAPI) ListClusterRoleBind(g *gin.Context) {
	list, err := w.clusterRoleBinding.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	clusterRoleBindingList := &rbacv1.ClusterRoleBindingList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, clusterRoleBindingList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, clusterRoleBindingList)
}
