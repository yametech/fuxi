package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	v1 "k8s.io/api/rbac/v1"
	"net/http"
)

// Get Role
func (w *WorkloadsAPI) GetRole(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.role.Get(dyn.ResourceRole, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Role
func (w *WorkloadsAPI) ListRole(g *gin.Context) {
	list, _ := w.role.List(dyn.ResourceRole, "", "", 0, 10000, nil)
	roleList := &v1.RoleList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, roleList)
	g.JSON(http.StatusOK, roleList)
}
