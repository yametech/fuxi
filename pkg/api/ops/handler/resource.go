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
		logging.Log.Error("create or update pipeline resource bind json error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "create or update pipeline resource  error:" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	if err := o.Service.CreateOrUpdatePipelineResource(rs); err != nil {
		logging.Log.Error("create or update pipeline resource error: " + err.Error())
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
	if namespace == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get task list error: namespace cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	pipelineResources, err := o.Service.PipelineResourceList(namespace)
	if err != nil {
		logging.Log.Error("get pipeline resource list error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get pipeline resource list  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "get pipeline resource list success",
		"code": http.StatusCreated,
		"data": pipelineResources,
	})
}

//PipelineResourceDelete deletes a pipeline resource
func (o *OpsController) PipelineResourceDelete(c *gin.Context) {

	namespace := c.Param("namespace")
	name := c.Param("name")
	if namespace == "" && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get task list error: namespace or name cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	err := o.Service.PipelineResourceDelete(namespace, name)
	if err != nil {
		logging.Log.Error("pipeline resource delete error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "pipeline resource delete error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "pipeline resource delete success",
		"code": http.StatusOK,
		"data": "",
	})
}

//GetPipelineResource get pipeline resource
func (o *OpsController) GetPipelineResource(c *gin.Context) {

	namespace := c.Param("namespace")
	name := c.Param("name")
	if namespace == "" && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get task list error: namespace or name cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	task, err := o.Service.GetPipelineResource(namespace, name)
	if err != nil {
		logging.Log.Error("get pipeline resource error:" + err.Error())
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
