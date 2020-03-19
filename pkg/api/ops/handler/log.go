package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"net/http"
)

//GetTaskRunLog get task run log
func (o *OpsController) GetTaskRunLog(c *gin.Context) {
	userName := o.getUserName(c)
	if userName == "" {
		return
	}

	namespace := c.Param("namespace")
	taskRunLog, err := o.Service.GetTaskRunLog(userName, namespace)

	if err != nil {
		logging.Log.Error("---------->GetTaskRunLog error: " + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get taskRunLog error" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get taskRunLog success",
		"code": http.StatusOK,
		"data": taskRunLog,
	})
}

//GetPipelineRunLog get pipeline run log
func (o *OpsController) GetPipelineRunLog(c *gin.Context) {
	userName := o.getUserName(c)
	if userName == "" {
		return
	}

	namespace := c.Param("namespace")
	pipelineRunLogs, err := o.Service.GetPipelineRunLog(userName, namespace)

	if err != nil {
		logging.Log.Error("---------->GetPipelineRunLog error:", err.Error())
		c.JSON(http.StatusPartialContent, gin.H{
			"msg":  "get pipelineRunLogs error" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get pipelineRunLogs success",
		"code": http.StatusOK,
		"data": pipelineRunLogs,
	})
}
