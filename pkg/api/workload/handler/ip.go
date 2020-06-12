package handler

import (
	"encoding/json"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"

	v1 "github.com/alauda/kube-ovn/pkg/apis/kubeovn/v1"
	"github.com/gin-gonic/gin"
)

// Get IP
func (w *WorkloadsAPI) GetIP(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.ip.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List IP
func (w *WorkloadsAPI) ListIP(g *gin.Context) {
	list, err := resourceList(g, w.ip)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	roleList := &v1.IPList{}
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
