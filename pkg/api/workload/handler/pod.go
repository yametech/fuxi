package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/workload/template"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
)

// Get Pod
func (w *WorkloadsAPI) GetPod(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.pod.Get(dyn.ResourcePod, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Pods
func (w *WorkloadsAPI) ListPod(g *gin.Context) {
	list, _ := w.pod.List(dyn.ResourcePod, "", "", 0, 10000, nil)
	podList := &corev1.PodList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{code: http.StatusBadRequest, data: "", msg: err.Error(), status: "Request bad parameter"})
		return
	}
	_ = json.Unmarshal(marshalData, podList)
	g.JSON(http.StatusOK, podList)
}

// AttachPod request and backend pod pty bing
func (w *WorkloadsAPI) AttachPod(g *gin.Context) {
	attachPodRequest := &template.AttachPodRequest{}
	attachPodRequest.Namespace = g.Param("namespace")
	attachPodRequest.Name = g.Param("name")
	attachPodRequest.Container = g.Param("container")

	sessionId, _ := GenerateTerminalSessionId()
	sharedSessionManager.set(sessionId,
		&SessionChannel{
			id:       sessionId,
			bound:    make(chan struct{}),
			sizeChan: make(chan remotecommand.TerminalSize),
		})

	go WaitForTerminal(attachPodRequest, sessionId)
	g.JSON(http.StatusOK, gin.H{"op": BIND, "sessionId": sessionId})
}
