package handler

import (
	"encoding/json"
	"fmt"
	"github.com/yametech/fuxi/pkg/api/workload/template"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
)

const END_OF_TRANSMISSION = "\u0004"

type SessionManager struct {
	restClient rest.RESTClient
	cfg        *rest.Config
}

// Process executed cmd in the container specified in request and connects it up with the  SessionChannel (a session)
func (sm *SessionManager) Process(
	request *template.AttachPodRequest,
	cmd []string,
	sess *SessionChannel,
) error {
	podExecOpt := &v1.PodExecOptions{
		Container: request.ContainerName,
		Command:   cmd,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}
	req := sm.restClient.Post().Resource("pods").
		Name(request.PodName).
		Namespace(request.Namespace).
		SubResource("exec").
		VersionedParams(podExecOpt, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(sm.cfg, "POST", req.URL())
	if err != nil {
		return err
	}
	return exec.Stream(remotecommand.StreamOptions{
		Stdin:             sess,
		Stdout:            sess,
		Stderr:            sess,
		TerminalSizeQueue: sess,
		Tty:               true,
	})
}

type SessionChannel struct {
	id            string
	bound         chan error
	sockJSSession sockjs.Session
	sizeChan      chan remotecommand.TerminalSize
	doneChan      chan struct{}
}

type TerminalMessage struct {
	Op, Data, SessionID string
	Rows, Cols          uint16
}

func (s *SessionChannel) Next() *remotecommand.TerminalSize {
	select {
	case size := <-s.sizeChan:
		return &size
	case <-s.doneChan:
		return nil
	}
}

func (s *SessionChannel) Write(p []byte) (int, error) {
	msg, err := json.Marshal(TerminalMessage{
		Op:   "stdout",
		Data: string(p),
	})
	if err != nil {
		return 0, err
	}

	if err = s.sockJSSession.Send(string(msg)); err != nil {
		return 0, err
	}
	return len(data), nil
}

func (s *SessionChannel) Read(p []byte) (n int, err error) {
	m, err := s.sockJSSession.Recv()
	if err != nil {
		return copy(p, END_OF_TRANSMISSION), err
	}

	var msg TerminalMessage
	if err := json.Unmarshal([]byte(m), &msg); err != nil {
		return copy(p, END_OF_TRANSMISSION), err
	}

	switch msg.Op {
	case "stdin":
		return copy(p, msg.Data), nil
	case "resize":
		s.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	default:
		return copy(p, END_OF_TRANSMISSION), fmt.Errorf("unknown message type '%s'", msg.Op)
	}
}

func CreateAttachHandler(prefix string) http.Handler {
	return sockjs.NewHandler(prefix, sockjs.DefaultOptions, func(session sockjs.Session) {

	})
}
