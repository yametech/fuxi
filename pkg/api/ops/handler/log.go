package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yametech/fuxi/pkg/logging"
	"net/http"
)

//GetTaskRunLog get task run log
func (o *OpsController) GetTaskRunLog(ctx *gin.Context) {

	check, namespace, name := o.CheckParams(ctx)
	if check {
		return
	}

	taskRunLog, err := o.Service.GetTaskRunLog(name, namespace)

	if err != nil {
		logging.Log.Error("get task run log error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get task run log error" + err.Error(),
			"code": http.StatusBadRequest,
			"data": "",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "get task run log success",
		"code": http.StatusOK,
		"data": taskRunLog,
	})
}

//GetPipelineRunLog get pipeline run log
func (o *OpsController) GetPipelineRunLog(c *gin.Context) {

	check, namespace, name := o.CheckParams(c)
	if check {
		return
	}
	pipelineRunLogs, err := o.Service.GetPipelineRunLog(name, namespace)

	if err != nil {
		logging.Log.Error("get pipeline run log error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get pipeline run logs error" + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "get pipeline run logs success",
		"code": http.StatusOK,
		"data": pipelineRunLogs,
	})
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (o *OpsController) GetRealLog(ctx *gin.Context) {

	check, namespace, name := o.CheckParams(ctx)
	if check {
		return
	}

	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get real log error: " + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
		return
	}

	defer ws.Close()

	mt, _, err := ws.ReadMessage()
	if err != nil {
		return
	}

	logC, errC, err := o.Service.ReadLivePipelineLogs(name, namespace, nil)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "get real log error: " + err.Error(),
			"code": http.StatusInternalServerError,
			"data": "",
		})
		return
	}

	for logC != nil || errC != nil {
		select {
		case l, ok := <-logC:
			if !ok {
				logC = nil
				continue
			}

			if l.Log == "EOFLOG" {
				continue
			}
			j, err := json.Marshal(l)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg":  "get real log error: " + err.Error(),
					"code": http.StatusInternalServerError,
					"data": "",
				})
			}

			ws.WriteMessage(mt, j)

		case e, ok := <-errC:
			if !ok {
				errC = nil
				continue
			}
			ws.WriteMessage(mt, []byte(e.Error()))
		}
	}

}
