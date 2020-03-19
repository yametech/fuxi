package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ops"
	"net/http"
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
		"msg":  "create or update pipeline run success",
		"code": http.StatusCreated,
		"data": "",
	})
}

//GetLatestPipelineRunList gets pipeline run list
func (o *OpsController) GetLatestPipelineRunList(c *gin.Context) {

	namespace := c.Param("namespace")
	paramsMap := make(map[string]string)
	paramsMap["namespace"] = namespace
	paramsMap["latest"] = "true"
	prs, err := o.Service.GetLatestPipelineRunList(namespace, paramsMap)

	if err != nil {
		logging.Log.Error("---------->PipelineRunList error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get latest pipeline run list error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get task run list success",
		"code": http.StatusOK,
		"data": prs,
	})
}

//GetPipelineRunHistoryList gets pipeline run list
func (o *OpsController) GetPipelineRunHistoryList(c *gin.Context) {

	//name: "ABC-1" will replace "ABC"
	name := c.Param("name")
	namespace := c.Param("namespace")
	paramsMap := make(map[string]string)
	paramsMap["name"] = name
	paramsMap["namespace"] = namespace
	paramsMap["latest"] = "true"
	prs, err := o.Service.GetPipelineRunList(namespace, paramsMap)

	if err != nil {
		logging.Log.Error("---------->GetPipelineRunHistoryList error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get  pipeline run history list error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get  pipeline run history list success",
		"code": http.StatusOK,
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
