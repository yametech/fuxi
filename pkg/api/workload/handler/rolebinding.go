package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/rbac/v1"
	"net/http"
)

// Get RoleBinding
func (w *WorkloadsAPI) GetRoleBinding(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.roleBinding.Get(dyn.ResourceRoleBinding, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List RoleBinding
func (w *WorkloadsAPI) ListRoleBinding(g *gin.Context) {
	list, _ := w.roleBinding.List(dyn.ResourceRoleBinding, "", "", 0, 10000, nil)
	roleBindingList := &v1.RoleBindingList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, roleBindingList)
	g.JSON(http.StatusOK, roleBindingList)
}
