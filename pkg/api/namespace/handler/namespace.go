package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ns"
	"net/http"
)

type NSController struct {
	Service ns.NS
}


//CreateSubNet create a subnet in ovn controller
func (n *NSController)CreateSubNet(c *gin.Context)  {
	var subnet ns.SubNet
	if err := c.ShouldBindJSON(&subnet); err != nil {
		logging.Log.Error("create subnet bind json error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "create subnet error:" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	if err :=n.Service.CreateSubnet(subnet); err != nil {
		logging.Log.Error("create subnet error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "create subnet error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "create subnet success",
		"code": http.StatusCreated,
		"data": "",
	})
}

//UpdateSubNet update a subnet in ovn controller
func (n *NSController) UpdateSubNet(c *gin.Context)  {
	var subnet ns.SubNet
	if err := c.ShouldBindJSON(&subnet); err != nil {
		logging.Log.Error("update subnet bind json error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "update subnet  error:" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	if err :=n.Service.CreateSubnet(subnet); err != nil {
		logging.Log.Error("update subnet error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "update subnet error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "update subnet success",
		"code": http.StatusOK,
		"data": "",
	})
}

//SubNetList get all subnets
func (n *NSController) SubNetList(c *gin.Context) {

	subnets, err := n.Service.SubNetList()
	if err != nil {
		logging.Log.Error("subnet list error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "subnet  list  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "subnet list success",
		"code": http.StatusOK,
		"data": subnets,
	})
}

//GetSubNet get a subnet
func (n *NSController) GetSubNet(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		logging.Log.Error("get subnet error: name param empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "et subnet error: name param empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	subnet, err := n.Service.GetSubNet(name)
	if err != nil {
		logging.Log.Error("get subnet error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get subnet  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get subnet  success",
		"code": http.StatusOK,
		"data": subnet,
	})
}

//DeleteSubNet deletes a subnet
func (n *NSController) DeleteSubNet(c *gin.Context) {
	name := c.Param("name")
	if len(name) == 0 {
		logging.Log.Error("delete subnet error: name param empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "delete subnet error: name param empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	err := n.Service.SubNetDelete(name)
	if err != nil {
		logging.Log.Error("delete subnet error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "delete subnet  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"msg":  "delete subnet  success",
		"code": http.StatusNoContent,
		"data": "",
	})
}