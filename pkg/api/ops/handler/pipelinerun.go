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
		logging.Log.Error("create or update pipelinerun bind json error:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "create or update pipelinerun  error:" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	if err := o.Service.CreateOrUpdatePipelineRun(&pr); err != nil {
		logging.Log.Error("create or update pipelinerun error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "create or update pipelinerun error:" + err.Error(),
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

//GetLatestPipelineRunList gets pipeline run list
func (o *OpsController) GetLatestPipelineRunList(c *gin.Context) {

	namespace := c.Param("namespace")
	paramsMap := make(map[string]string)
	paramsMap["namespace"] = namespace
	paramsMap["latest"] = "true"
	prs, err := o.Service.GetLatestPipelineRunList(namespace, paramsMap)

	if err != nil {
		logging.Log.Error("get latest pipelinerun list error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get latest pipelinerun list error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get pipelinerun list success",
		"code": http.StatusOK,
		"data": prs,
	})
}

//GetPipelineRunHistoryList gets pipeline run list
func (o *OpsController) GetPipelineRunHistoryList(c *gin.Context) {

	name := c.Param("name")
	namespace := c.Param("namespace")
	if namespace == "" && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get real log error: namespace or name cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	//name: "ABC-1" will replace "ABC"
	paramsMap := make(map[string]string)
	paramsMap["name"] = name
	paramsMap["namespace"] = namespace
	paramsMap["latest"] = "true"
	prs, err := o.Service.GetPipelineRunList(namespace, paramsMap)

	if err != nil {
		logging.Log.Error("get pipelinerun history list error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get  pipelinerun history list error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get  pipelinerun history list success",
		"code": http.StatusOK,
		"data": prs,
	})
}

//PipelineRunDelete delete a pipeline run
func (o *OpsController) PipelineRunDelete(c *gin.Context) {

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

	err := o.Service.PipelineRunDelete(namespace, name)
	if err != nil {
		logging.Log.Error("pipelinerun delete error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "pipelinerun delete error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusNoContent, gin.H{
		"msg":  "pipelinerun delete success",
		"code": http.StatusNoContent,
		"data": "",
	})
}

//GetPipelineRun get pipeline run
func (o *OpsController) GetPipelineRun(c *gin.Context) {

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

	p, err := o.Service.GetPipelineRun(name, namespace)
	if err != nil {
		logging.Log.Error("get pipelinerun error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get pipelinerun error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get pipelinerun success",
		"code": http.StatusOK,
		"data": p,
	})
}

//ReRunPipeline rerun a pipeline run
func (o *OpsController) ReRunPipeline(c *gin.Context) {

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

	err := o.Service.ReRunPipeline(name, namespace)
	if err != nil {
		logging.Log.Error("rerun pipeline error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "rerun pipelinerun  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "rerun  pipelinerun success",
		"code": http.StatusOK,
		"data": "",
	})
}

//CancelPipelineRun cancel a pipeline run
func (o *OpsController) CancelPipelineRun(c *gin.Context) {

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

	err := o.Service.CancelPipelineRun(namespace, name)
	if err != nil {
		logging.Log.Error("cancel pipelinerun error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "cancel pipelinerun error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusNoContent, gin.H{
		"msg":  "",
		"code": http.StatusNoContent,
		"data": "",
	})
}
