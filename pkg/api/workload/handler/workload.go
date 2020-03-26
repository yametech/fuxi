package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/workload/template"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"net/http"
)

type WorkloadsApi struct{}

func (w *WorkloadsApi) ListStone(g *gin.Context) {
	stone := &template.StoneRequest{}
	if err := g.ShouldBind(stone); err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter"},
		)
		return
	}
}

func (w *WorkloadsApi) PodAttach(g *gin.Context) {
	attachPodRequest := &template.AttachPodRequest{}
	if err := g.Bind(attachPodRequest); err != nil {
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
	sessionId, _ := GenTerminalSessionId()
	sharedSessionManager.set(sessionId, &SessionChannel{
		id:       sessionId,
		bound:    make(chan error),
		sizeChan: make(chan remotecommand.TerminalSize),
	})

	log.Printf("attach pod (%s).(%s) sessionId (%s)\n", attachPodRequest.Namespace, attachPodRequest.PodName, sessionId)

	go WaitForTerminal(attachPodRequest, sessionId)
	g.JSON(http.StatusOK, gin.H{"op": BIND, "session_id": sessionId})
}
