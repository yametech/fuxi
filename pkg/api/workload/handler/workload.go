package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/workload/template"
	"net/http"
)

type WorkloadsApi struct{}

func (w *WorkloadsApi) ListStone(g *gin.Context) {
	stone := &template.StoneRequest{}
	if err := g.ShouldBind(stone); err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter"},
		)
		return
	}
}


//func (w *WorkloadApi) ListDeployments(c *gin.Context) {
//	deploy := &template.Deployment{}
//	if err := c.ShouldBind(deploy); err != nil {
//		//gin.H{}
//	}
//
//	//
//}
