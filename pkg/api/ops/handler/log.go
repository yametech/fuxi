package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yametech/fuxi/pkg/logging"
	"golang.org/x/sync/errgroup"
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


var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}


func(o *OpsController) GetRealLog(ctx *gin.Context){

	namespace := ctx.Param("namespace")
	name := ctx.Param("name")

	if namespace == "" && name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get GetRealLog error: namespace or name cannot be empty",
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "get GetRealLog error: " + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
		return
	}

	defer ws.Close()


	mt, _, err := ws.ReadMessage()
	if err != nil {
		return
	}

	var logs = make(chan []string,10)
	var g errgroup.Group
	g.Go(func() error{
		err := o.Service.GetTaskRealLog(name,namespace,logs)
		if err != nil {
			close(logs)
		}
		return err
	})

	if err := g.Wait(); err == nil {
		for{

			select{

			case ss,ok:=<- logs:
				if !ok {
					return
				}

				for i :=range ss {
					fmt.Println(ss[i])
					ws.WriteMessage(mt, []byte(ss[i]))
				}

			}
		}
	}


}

