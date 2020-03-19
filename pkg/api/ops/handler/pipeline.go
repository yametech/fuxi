package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ops"
	"net/http"
)

//CreateOrUpdatePipeline creates or update pipeline
func (o *OpsController) CreateOrUpdatePipeline(c *gin.Context) {

	var pipeline ops.Pipeline
	if err := c.ShouldBindJSON(&pipeline); err != nil {
		logging.Log.Error("---------->CreateOrUpdatePipeline should bind json error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "CreateOrUpdatePipeline  error:" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	if err := o.Service.CreateOrUpdatePipeline(pipeline); err != nil {
		logging.Log.Error("---------->CreateOrUpdatePipeline error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "create or update pipeline" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "create or update pipeline",
		"code": http.StatusCreated,
		"data": "",
	})
}

// PipelineList returns a list of pipelines
func (o *OpsController) PipelineList(c *gin.Context) {
	namespace := c.Param("namespace")
	pipelines, err := o.Service.PipelineList(namespace)
	if err != nil {
		logging.Log.Error("---------->PipelineList error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get  pipelines error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "get taskRunLog success",
		"code": http.StatusCreated,
		"data": pipelines,
	})
}

//PipelineDelete deletes a pipeline
func (o *OpsController) PipelineDelete(c *gin.Context) {
	userName := o.getUserName(c)
	if userName == "" {
		return
	}
	namespace := c.Param("namespace")
	err := o.Service.PipelineDelete(userName, namespace)

	if err != nil {
		logging.Log.Error("---------->PipelineDelete error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "delete pipeline error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "delete pipeline success",
		"code": http.StatusOK,
		"data": "",
	})
}

//GetPipeline get pipeline
func (o *OpsController) GetPipeline(c *gin.Context) {

	userName := o.getUserName(c)
	if userName == "" {
		return
	}
	namespace := c.Param("namespace")
	p, err := o.Service.GetPipeline(userName, namespace)
	if err != nil {
		logging.Log.Error("---------->GetPipeline error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get pipeline error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get pipeline success",
		"code": http.StatusOK,
		"data": p,
	})
}
