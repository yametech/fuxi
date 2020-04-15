package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	"net/http"
	"net/url"
	"os"
)

var event = workloadservice.NewEvent()

type MyEvent struct {
	Type   watch.EventType `json:"type"`
	Object runtime.Object  `json:"object"`
}

func watchData(g *gin.Context) {

	refererUrl := g.Request.Header.Get("Referer")
	u, err := url.Parse(refererUrl)
	_ = err
	//if err != nil {
	//	g.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//}

	var resourceType schema.GroupVersionResource

	switch u.RawPath {
	case "events":
		resourceType = dyn.ResourceEvent
	case "pods":
		resourceType = dyn.ResourcePod
	default:
		resourceType = dyn.ResourcePod
	}

	fmt.Fprintf(os.Stdout, "watch come from %s\r\n", refererUrl)

	w := g.Writer
	header := w.Header()
	header.Set("Transfer-Encoding", "chunked")
	header.Set("Content-Type", "text/event-stream")
	header.Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	closed := make(chan struct{})
	item, _ := event.Watch(resourceType, "", nil, closed)

	//after := time.After(1 * time.Second)
	for {
		select {
		//case <-after:
		//	closed <- struct{}{}
		//	w.WriteString("data: STREAM_END \n\n")
		case obj, ok := <-item:
			if !ok {
				return
			}
			myevent := &MyEvent{Object: obj.Object, Type: obj.Type}
			b, err := json.Marshal(myevent)
			if err != nil {
				return
			}
			_, err = w.WriteString(fmt.Sprintf(`data: %s%s`, string(b), "\n\n"))
			if err != nil {
				return
			}
			w.(http.Flusher).Flush()
		}
	}
}
