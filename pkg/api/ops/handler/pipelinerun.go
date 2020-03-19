package handler

import (
<<<<<<< HEAD
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ops"
	"net/http"
=======
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ops"
>>>>>>> 0c71acc6e0202644d124b914f3c302d8c1d93ea5
)

//CreateOrUpdatePipelineRun creates or updates a pipeline run
func (o *OpsController) CreateOrUpdatePipelineRun(c *gin.Context) {

	var pr ops.PipelineRun
	if err := c.ShouldBindJSON(&pr); err != nil {
		logging.Log.Error("---------->CreateOrUpdatePipelineRun bind json error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "CreateOrUpdatePipelineRun  error:" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	if err := o.Service.CreateOrUpdatePipelineRun(&pr); err != nil {
		logging.Log.Error("---------->CreateOrUpdatePipelineRun error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "create or update pipelinerun" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "create or update pipelinerun success",
		"code": http.StatusCreated,
		"data": "",
	})
}

//PipelineRunList gets pipeline run list
func (o *OpsController) PipelineRunList(c *gin.Context) {

	namespace := c.Param("namespace")
	prs, err := o.Service.PipelineRunList(namespace)

	if err != nil {
		logging.Log.Error("---------->PipelineRunList error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get  pipelines error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "get task run list success",
		"code": http.StatusCreated,
		"data": prs,
	})
}

//PipelineRunDelete delete a pipeline run
func (o *OpsController) PipelineRunDelete(c *gin.Context) {

	userName := o.getUserName(c)
	if userName == "" {
		return
	}
	namespace := c.Param("namespace")
	err := o.Service.PipelineRunDelete(userName, namespace)
	if err != nil {
		logging.Log.Error("---------->PipelineRunDelete error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "delete pipeline run error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "delete pipeline run success",
		"code": http.StatusOK,
		"data": "",
	})
}

//GetPipelineRun get pipeline run
func (o *OpsController) GetPipelineRun(c *gin.Context) {

	userName := o.getUserName(c)
	if userName == "" {
		return
	}
	namespace := c.Param("namespace")
	p, err := o.Service.GetPipelineRun(userName, namespace)
	if err != nil {
		logging.Log.Error("---------->GetPipelineRun error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get pipeline run error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get pipeline run success",
		"code": http.StatusOK,
		"data": p,
	})
}

//ReRunPipeline rerun a pipeline run
func (o *OpsController) ReRunPipeline(c *gin.Context) {

	userName := o.getUserName(c)
	if userName == "" {
		return
	}
	namespace := c.Param("namespace")
	err := o.Service.ReRunPipeline(userName, namespace)
	if err != nil {
		logging.Log.Error("---------->ReRunPipeline error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "ReRun Pipeline run  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "rerun  pipeline run success",
		"code": http.StatusOK,
		"data": "",
	})
}
