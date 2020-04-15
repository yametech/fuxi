package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/logging"
	"github.com/yametech/fuxi/pkg/service/ops"
	"net/http"
)

//CreateOrUpdateTask creates or update task
func (o *OpsController) CreateOrUpdateTask(c *gin.Context) {

	var task ops.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		logging.Log.Error("create or update task bind json error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "create or update task  error:" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	if err := o.Service.CreateOrUpdateTask(&task); err != nil {
		logging.Log.Error("create or update task error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "create or update task  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "create or update task success",
		"code": http.StatusCreated,
		"data": "",
	})
}

//TaskList get task list
func (o *OpsController) TaskList(c *gin.Context) {

	namespace := c.Param("namespace")
	if namespace == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get task list error: namespace cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	tasks, err := o.Service.TaskList(namespace)
	if err != nil {
		logging.Log.Error("task list error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get  task  error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "get task list success",
		"code": http.StatusCreated,
		"data": tasks,
	})
}

//DeleteTask delete a task
func (o *OpsController) DeleteTask(c *gin.Context) {

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

	err := o.Service.DeleteTask(name, namespace)
	if err != nil {
		logging.Log.Error("delete task error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "delete task error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "delete task success",
		"code": http.StatusOK,
		"data": "",
	})
}

//GetTask get the task
func (o *OpsController) GetTask(c *gin.Context) {

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

	task, err := o.Service.GetTask(name, namespace)
	if err != nil {
		logging.Log.Error("get task error: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get task error:" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get task success",
		"code": http.StatusOK,
		"data": task,
	})
}
