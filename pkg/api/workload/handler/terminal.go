/*
// terminal based kubernetes/dashboard terminal implement ideas, thanks for kubernetes authors, great open source contributor
// Copyright 2020 The yametech Authors.
// Copyright 2017 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

*/

package handler

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/client-go/kubernetes"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/yametech/fuxi/pkg/api/workload/template"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

// END_OF_TRANSMISSION terminal end of
const END_OF_TRANSMISSION = "\u0004"

// OP      DIRECTION  FIELD(S) USED  DESCRIPTION
type OP uint8

const (
	// BIND    fe->be     sessionID      Id sent back from TerminalResponse
	BIND = iota // 0
	// STDIN   fe->be     Data           Keystrokes/paste buffer
	STDIN
	// STDOUT  be->fe     Data           Output from the process
	STDOUT
	// RESIZE  fe->be     Rows, Cols     New terminal size
	RESIZE
	// TOAST   be->fe     Data           OOB message to be shown to the user
	TOAST
)

// Global import the package init the session manager
var sharedSessionManager *sessionManager

// CreateSharedSessionManager none
func CreateSharedSessionManager(clientSet *kubernetes.Clientset, restCfg *rest.Config) {
	if sharedSessionManager == nil {
		sharedSessionManager = &sessionManager{
			client:   clientSet.CoreV1().RESTClient(),
			restCfg:  restCfg,
			channels: make(map[string]*SessionChannel),
			lock:     sync.RWMutex{},
		}
	}
	return
}

// sessionManager manager external client connect to kubernetes pod terminal session
type sessionManager struct {
	client   rest.Interface
	restCfg  *rest.Config
	channels map[string]*SessionChannel
	lock     sync.RWMutex
}

func (sm *sessionManager) get(sessionId string) (*SessionChannel, bool) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	v, exists := sm.channels[sessionId]
	return v, exists

}

func (sm *sessionManager) set(sessionId string, channel *SessionChannel) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sm.channels[sessionId] = channel
}

// close shuts down the SockJS connection and sends the status code and reason to the client
// Can happen if the process exits or if there is an error starting up the process
// For now the status code is unused and reason is shown to the user (unless "")
func (sm *sessionManager) close(sessionID string, status uint32, reason string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	err := sm.channels[sessionID].sockJSSession.Close(status, reason)
	if err != nil {
		return
	}
	delete(sm.channels, sessionID)
}

// PtyHandler is what remotecommand expects from a pty
type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

// process executed cmd in the container specified in request and connects it up with the  SessionChannel (a session)
func (sm *sessionManager) process(request *template.AttachPodRequest, cmd []string, sess PtyHandler) error {
	command := []string{"/bin/sh", "-c"}
	command = append(command, cmd...)
	req := sm.client.Post().
		Resource("pods").
		Name(request.Name).
		Namespace(request.Namespace).
		SubResource("exec")
	option := &v1.PodExecOptions{
		Command: command,
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}
	if request.Container != "" {
		option.Container = request.Container
	}
	req.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(sm.restCfg, "POST", req.URL())
	if err != nil {
		return err
	}
	err = exec.Stream(
		remotecommand.StreamOptions{
			Stdin:             sess,
			Stdout:            sess,
			Stderr:            sess,
			TerminalSizeQueue: sess,
			Tty:               true,
		})

	if err != nil {
		return err
	}
	return nil
}

// SessionChannel a http connect
// upgrade to websocket session bind a session channel to backend kubernetes API server with SPDY
type SessionChannel struct {
	id            string
	bound         chan struct{}
	sockJSSession sockjs.Session
	sizeChan      chan remotecommand.TerminalSize
	doneChan      chan struct{}
}

// TerminalMessage is the messaging protocol between ShellController and TerminalSession.
type TerminalMessage struct {
	Data, SessionID string
	Rows, Cols      uint16
	Op              OP
}

// Next impl sizeChan remotecommand.TerminalSize
func (s *SessionChannel) Next() *remotecommand.TerminalSize {
	select {
	case size := <-s.sizeChan:
		return &size
	case <-s.doneChan:
		return nil
	}
}

// Write impl io.Writer
func (s *SessionChannel) Write(p []byte) (int, error) {
	msg, err := json.Marshal(TerminalMessage{
		Op:   STDOUT,
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

// Toast can be used to send the user any OOB messages
// hterm puts these in the center of the terminal
func (s *SessionChannel) Toast(p string) error {
	msg, err := json.Marshal(TerminalMessage{
		Op:   TOAST,
		Data: p,
	})
	if err != nil {
		return err
	}

	if err = s.sockJSSession.Send(string(msg)); err != nil {
		return err
	}
	return nil
}

// Read impl io.Reader
func (s *SessionChannel) Read(p []byte) (n int, err error) {
	//retry:
	//	m, err := s.sockJSSession.Recv()
	//	if err != nil {
	//		return copy(p, END_OF_TRANSMISSION), err
	//	}
	//
	//	op, err := base64.RawStdEncoding.DecodeString(m)
	//	if err != nil {
	//		goto retry
	//	}
	//	if len(op) < 1 {
	//		goto retry
	//	}
	//	switch op[0] {
	//	case '(':
	//		goto retry
	//	case ')':
	//		goto retry
	//	}
	m, err := s.sockJSSession.Recv()
	var msg TerminalMessage
	if err := json.Unmarshal([]byte(m), &msg); err != nil {
		//goto retry
		return copy(p, END_OF_TRANSMISSION), err
	}

	switch msg.Op {
	case STDIN:
		return copy(p, msg.Data), nil
	case RESIZE:
		s.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	default:
		return copy(p, END_OF_TRANSMISSION), fmt.Errorf("unknown message type '%d'", msg.Op)
	}
}

// CreateAttachHandler is called from main for /workload/attach
func CreateAttachHandler(path string) http.Handler {
	return sockjs.NewHandler(path, sockjs.DefaultOptions, HandleTerminalSession)
}

func HandleTerminalSession(session sockjs.Session) {
	_, _ = session.Recv()
	buf, err := session.Recv()
	if err != nil {
		log.Printf("recv buffer error: %s", err)
		return
	}

	msg := &TerminalMessage{}
	if err = json.Unmarshal([]byte(buf), &msg); err != nil {
		log.Printf("handleTerminalSession: can't un marshal (%v): %s", buf, err)
		return
	}
	if msg.Op != BIND {
		log.Printf("handleTerminalSession: expected 'bind' message, got: %s", buf)
		return
	}
	terminalSession, exist := sharedSessionManager.get(msg.SessionID)
	if !exist {
		log.Printf("sharedSessionManager: can't find session '%s'", msg.SessionID)
		return
	}
	if terminalSession.id == "" {
		log.Printf("handleTerminalSession: can't find session '%s'", msg.SessionID)
		return
	}
	terminalSession.sockJSSession = session
	sharedSessionManager.set(msg.SessionID, terminalSession)
	terminalSession.bound <- struct{}{}
}

// GenerateTerminalSessionId generates a random session ID string. The format is not really interesting.
// This ID is used to identify the session when the client opens the SockJS connection.
// Not the same as the SockJS session id! We can't use that as that is generated
// on the client side and we don't have it yet at this point.
func GenerateTerminalSessionId() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	id := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(id, bytes)
	return string(id), nil
}

// isValidShell checks if the shell is an allowed one
func isValidShell(validShells []string, shell string) bool {
	for _, validShell := range validShells {
		if validShell == shell {
			return true
		}
	}
	return false
}

// WaitForTerminal is called from pod attach api as a goroutine
// Waits for the SockJS connection to be opened by the client the session to be bound in handleTerminalSession
func WaitForTerminal(request *template.AttachPodRequest, sessionId string) {
	if request.Shell == "" {
		request.Shell = "sh"
	}
	session, exist := sharedSessionManager.get(sessionId)
	if !exist {
		//
	}
	<-session.bound

	defer close(session.bound)
	var err error
	validShells := []string{"sh", "csh", "bash"}

	if isValidShell(validShells, request.Shell) {
		cmd := []string{request.Shell}
		err = sharedSessionManager.process(request, cmd, session)
	} else {
		// No shell given or it was not valid: try some shells until one succeeds or all fail
		for _, testShell := range validShells {
			cmd := []string{testShell}
			if err = sharedSessionManager.process(request, cmd, session); err == nil {
				break
			}
		}
	}
	if err != nil {
		sharedSessionManager.close(sessionId, 2, err.Error())
		return
	}
	sharedSessionManager.close(sessionId, 1, "process exited")

}
