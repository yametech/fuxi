package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ops"
)

type OpsController struct {
	Service ops.OpsService
}

func (o *OpsController) getUserName(c *gin.Context) string {
	userName := c.Request.Header.Get("x-auth-username")
	if len(userName) < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
			"msg":  "authorization not correct",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return ""
	}
	return userName

}

//ListRepos return current user department all repos
func (o *OpsController) ListRepos(c *gin.Context) {

	userName := o.getUserName(c)
	if userName == "" {
		return
	}

	namespace := c.Param("namespace")
	repos, err := o.Service.ListRepos(userName, namespace)
	if err != nil {
		logging.Log.Error("---------->ListRepos error:" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get repos failed",
			"code": http.StatusBadRequest,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get repos success",
		"code": http.StatusOK,
		"data": repos,
	})
}

//ListBranchs return current user department  all branchs
func (o *OpsController) ListBranchs(c *gin.Context) {
	userName := o.getUserName(c)
	if userName == "" {
		return
	}

	namespace := c.Param("namespace")
	branchs, err := o.Service.ListBranchs(userName, namespace)

	if err != nil {
		logging.Log.Error("---------->ListBranchs error:" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get branch error" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get branch success",
		"code": http.StatusOK,
		"data": branchs,
	})
}
