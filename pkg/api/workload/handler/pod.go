package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
	"time"
)

type logRequest struct {
	Container  string    `form:"container" json:"container"`
	Timestamps bool      `form:"timestamps" json:"timestamps"`
	SinceTime  time.Time `form:"sinceTime" json:"sinceTime"`
	TailLines  int64     `form:"tailLines" json:"tailLines"`
}

func (w *WorkloadsAPI) LogPod(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	lq := &logRequest{}
	if err := g.Bind(lq); err != nil || namespace == "" || name == "" {
		common.ToRequestParamsError(g, err)
		return
	}

	buf := bytes.NewBufferString("")
	err := w.pod.Logs(
		namespace,
		name,
		lq.Container,
		false,
		false,
		lq.Timestamps,
		0,
		&lq.SinceTime,
		0,
		lq.TailLines,
		buf,
	)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	g.JSON(http.StatusOK, buf.String())
}

// Get Pod
func (w *WorkloadsAPI) GetPod(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.pod.Get(namespace, name)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Pods
func (w *WorkloadsAPI) ListPod(g *gin.Context) {
	list, err := resourceList(g, w.pod)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	podList := &corev1.PodList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	err = json.Unmarshal(marshalData, podList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, podList)
}

type AttachPodRequest struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Container string `json:"container"`
	Shell     string `json:"shell"`
}

// AttachPod request and backend pod pty bing
func (w *WorkloadsAPI) AttachPod(g *gin.Context) {
	attachPodRequest := &AttachPodRequest{
		Namespace: g.Param("namespace"),
		Name:      g.Param("name"),
		Container: g.Param("container"),
		Shell:     g.Query("shell"),
	}

	sessionId, _ := generateTerminalSessionId()
	sharedSessionManager.set(sessionId,
		&sessionChannels{
			id:       sessionId,
			bound:    make(chan struct{}),
			sizeChan: make(chan remotecommand.TerminalSize),
		})

	go waitForTerminal(attachPodRequest, sessionId)
	g.JSON(http.StatusOK, gin.H{"op": BIND, "sessionId": sessionId})
}
