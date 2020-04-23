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
	"time"
)

//GET /workload/api/v1/namespaces/:namespace/pods/:name/log?container=controller&timestamps=true&tailLines=1000&sinceTime=2020-04-23T02%3A13%3A16.273Z
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
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter"},
		)
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
		g.JSON(
			http.StatusInternalServerError,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter"},
		)
	}

	g.JSON(http.StatusOK, buf.String())
}

// Get Pod
func (w *WorkloadsAPI) GetPod(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.pod.Get(dyn.ResourcePod, namespace, name)
	if err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter"},
		)
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
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter"},
		)
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
	attachPodRequest.Shell = g.Query("shell")

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
