package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yametech/fuxi/pkg/logging"
	"net/http"
	"time"
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


var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}


func(o *OpsController) GetRealLog(ctx *gin.Context){
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		time.Sleep(2000)
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

