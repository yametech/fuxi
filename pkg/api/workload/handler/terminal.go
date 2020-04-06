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
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/yametech/fuxi/pkg/api/workload/template"
	"github.com/yametech/fuxi/pkg/k8s/client"
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
	// RESIZE  fe->be     Rows, Cols     New terminal size
	RESIZE
	// STDOUT  be->fe     Data           Output from the process
	STDOUT
	// TOAST   be->fe     Data           OOB message to be shown to the user
	TOAST
)

// Global import the package init the session manager
var sharedSessionManager *sessionManager

// CreateSharedSessionManager none
func CreateSharedSessionManager() {
	if sharedSessionManager == nil {
		sharedSessionManager = &sessionManager{
			client:   client.K8sClient.RESTClient(),
			cfg:      client.RestConf,
			channels: make(map[string]*SessionChannel),
			lock:     sync.RWMutex{},
		}
	}
	return
}

// sessionManager manager external client connect to kubernetes pod terminal session
type sessionManager struct {
	client   rest.Interface
	cfg      *rest.Config
	channels map[string]*SessionChannel
	lock     sync.RWMutex
}

func (sm *sessionManager) get(sessionID string) *SessionChannel {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	v, _ := sm.channels[sessionID]
	return v

}

func (sm *sessionManager) set(sessionID string, channel *SessionChannel) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sm.channels[sessionID] = channel
}

// close shuts down the SockJS connection and sends the status code and reason to the client
// Can happen if the process exits or if there is an error starting up the process
// For now the status code is unused and reason is shown to the user (unless "")
func (sm *sessionManager) close(sessionID string, status uint32, reason string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	err := sm.channels[sessionID].sockJSSession.Close(status, reason)
	if err != nil {
		log.Print(err)
		return
	}
	delete(sm.channels, sessionID)
}

// process executed cmd in the container specified in request and connects it up with the  SessionChannel (a session)
func (sm *sessionManager) process(request *template.AttachPodRequest, cmd []string, sess *SessionChannel) error {
	beautifyCmd := []string{
		"/bin/sh",
		"-c",
		`TERM=xterm-256color; export TERM; [ -x /bin/bash ] && ([ -x /usr/bin/script ] && /usr/bin/script -q -c "/bin/bash" /dev/null || exec /bin/bash) || exec /bin/sh`,
	}
	beautifyCmd = append(beautifyCmd, cmd...)

	podExecOpt := &v1.PodExecOptions{
		Container: request.ContainerName,
		Command:   beautifyCmd,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}
	req := sm.client.Post().Resource("pods").
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

// SessionChannel a http connect
// upgrade to websocket session bind a session channel to backend kubernetes API server with SPDY
type SessionChannel struct {
	id            string
	bound         chan error
	sockJSSession sockjs.Session
	sizeChan      chan remotecommand.TerminalSize
	doneChan      chan struct{}
}

// TerminalMessage is the messaging protocol between ShellController and TerminalSession.
type TerminalMessage struct {
	Data, sessionID string
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
	m, err := s.sockJSSession.Recv()
	if err != nil {
		return copy(p, END_OF_TRANSMISSION), err
	}

	var msg TerminalMessage
	if err := json.Unmarshal([]byte(m), &msg); err != nil {
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
	return sockjs.NewHandler(path, sockjs.DefaultOptions, handleTerminalSession)
}

func handleTerminalSession(session sockjs.Session) {
	buf, err := session.Recv()
	if err != nil {
		log.Printf("recv buffer error: %s", err)
		return
	}
	msg := &TerminalMessage{}
	if err = json.Unmarshal([]byte(buf), &msg); err != nil {
		log.Printf("handleTerminalSession: can't un marshal (%v): %s", err, buf)
		return
	}
	if msg.Op != BIND {
		log.Printf("handleTerminalSession: expected 'bind' message, got: %s", buf)
		return
	}
	terminalSession := sharedSessionManager.get(msg.sessionID)
	if terminalSession.id == "" {
		log.Printf("handleTerminalSession: can't find session '%s'", msg.sessionID)
		return
	}
	terminalSession.sockJSSession = session
	sharedSessionManager.set(msg.sessionID, terminalSession)
	terminalSession.bound <- nil
}

// GenTerminalsessionID generates a random session ID string. The format is not really interesting.
// This ID is used to identify the session when the client opens the SockJS connection.
// Not the same as the SockJS session id! We can't use that as that is generated
// on the client side and we don't have it yet at this point.
func GenTerminalsessionID() (string, error) {
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
func WaitForTerminal(request *template.AttachPodRequest, sessionID string) {
	if request.Shell == "" {
		request.Shell = "sh"
	}
	select {
	case <-sharedSessionManager.get(sessionID).bound:
		defer close(sharedSessionManager.get(sessionID).bound)
		var err error
		validShells := []string{"bash", "sh", "powershell", "cmd"}

		if isValidShell(validShells, request.Shell) {
			cmd := []string{request.Shell}
			err = sharedSessionManager.process(request, cmd, sharedSessionManager.get(sessionID))
		} else {
			// No shell given or it was not valid: try some shells until one succeeds or all fail
			for _, testShell := range validShells {
				cmd := []string{testShell}
				if err = sharedSessionManager.process(request, cmd, sharedSessionManager.get(sessionID)); err == nil {
					break
				}
			}
		}
		if err != nil {
			sharedSessionManager.close(sessionID, 2, err.Error())
			return
		}
		sharedSessionManager.close(sessionID, 1, "process exited")
	}
}
