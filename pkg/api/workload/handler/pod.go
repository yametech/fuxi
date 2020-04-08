package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/workload/template"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	workload "github.com/yametech/fuxi/pkg/service/workload"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
)

//GetPod return the pod info allow list all
func (w *WorkloadsAPI) GetPod(g *gin.Context) {
	podRequest := &template.PodRequest{}
	if err := g.ShouldBind(podRequest); err != nil {
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
	podService := workload.NewPod()
	item, err := podService.Get(dyn.ResourcePod, *podRequest.Namespace, podRequest.Name)
	if err != nil {
		g.JSON(http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   "",
				msg:    err.Error(),
				status: "internal get pod error",
			},
		)
		return
	}
	g.JSON(http.StatusOK, item)
}

// ListPod list namespace pod, admin
func (w *WorkloadsAPI) ListPod(g *gin.Context) {
	//podList := corev1.PodList{}
}

// AttachPod request and backend pod pty bing
func (w *WorkloadsAPI) AttachPod(g *gin.Context) {
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
