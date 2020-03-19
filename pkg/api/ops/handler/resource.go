package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ops"
)

//CreateOrUpdatePipelineResource creates or update pipeline resource
func (o *OpsController) CreateOrUpdatePipelineResource(c *gin.Context) {

	var rs ops.Resource
	if err := c.ShouldBindJSON(&rs); err != nil {
		logging.Log.Error("---------->CreateOrUpdatePipelineResource bind json error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "CreateOrUpdatePipelineResource  error:" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	if err := o.Service.CreateOrUpdatePipelineResource(rs); err != nil {
		logging.Log.Error("---------->CreateOrUpdatePipelineResource error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "create or update pipeline resource" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "create or update pipeline resource success",
		"code": http.StatusCreated,
		"data": "",
	})
}

// PipelineResourceList get all pipeline resource
func (o *OpsController) PipelineResourceList(c *gin.Context) {

	namespace := c.Param("namespace")
	pipelineResources, err := o.Service.PipelineResourceList(namespace)
	if err != nil {
		logging.Log.Error("---------->PipelineResourceList error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get  pipeline resources  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "get pipeline resources success",
		"code": http.StatusCreated,
		"data": pipelineResources,
	})
}

//PipelineResourceDelete deletes a pipeline resource
func (o *OpsController) PipelineResourceDelete(c *gin.Context) {

	userName := o.getUserName(c)
	if userName == "" {
		return
	}
	namespace := c.Param("namespace")
	err := o.Service.PipelineResourceDelete(userName, namespace)
	if err != nil {
		logging.Log.Error("---------->PipelineResourceDelete error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "delete pipeline resource error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "delete pipeline resource success",
		"code": http.StatusOK,
		"data": "",
	})
}

//GetPipelineResource get pipeline resource
func (o *OpsController) GetPipelineResource(c *gin.Context) {

	userName := o.getUserName(c)
	if userName == "" {
		return
	}
	namespace := c.Param("namespace")
	task, err := o.Service.GetPipelineResource(userName, namespace)
	if err != nil {
		logging.Log.Error("---------->GetPipelineResource error:" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get pipeline resource error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get pipeline resource success",
		"code": http.StatusOK,
		"data": task,
	})
}
