package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/workload/template"
	"k8s.io/client-go/tools/remotecommand"
)

//GetPod return the pod info
func (w *WorkloadsAPI) GetPod(g *gin.Context) {
}

// PodAttach request and backend pod pty bing
func (w *WorkloadsAPI) PodAttach(g *gin.Context) {
	attachPodRequest := &template.AttachPodRequest{}
	if err := g.ShouldBind(attachPodRequest); err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter",
			},
		)
		return
	}
	sessionID, _ := GenTerminalsessionID()
	sharedSessionManager.set(sessionID, &SessionChannel{
		id:       sessionID,
		bound:    make(chan error),
		sizeChan: make(chan remotecommand.TerminalSize),
	})

	go WaitForTerminal(attachPodRequest, sessionID)
	g.JSON(http.StatusOK, gin.H{"op": BIND, "sessionId": sessionID})
}
