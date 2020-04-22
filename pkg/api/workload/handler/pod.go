package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/workload/template"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
)

func (w *WorkloadsAPI) LogPod(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	container := g.Query("container")
	//timestamps := g.Query("timestamps")
	//tailLines := g.Query("tailLines")
	//sinceTime := g.Query("sinceTime")
	buf := bytes.NewBufferString("")
	//date := time.Date(1970, 1, 0, 0, 0, 0, 0, &time.Location{})
	w.pod.Logs(namespace, name, container, false, false, true, 2000,
		nil, 1024000, 1000, buf)
	g.JSON(http.StatusOK, buf.String())
}

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
