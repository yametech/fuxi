package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Event struct {
	Type   watch.EventType `json:"type"`
	Object runtime.Object  `json:"object"`
	Url    string          `json:"url"`
	Status int             `json:"status"`
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func parseForApiUrl(apiUrl string) (gvr *schema.GroupVersionResource, namespace string, resourceVersion string, err error) {
	u, err := url.Parse(apiUrl)
	if err != nil {
		return nil, "", "", err
	}
	gvr = &schema.GroupVersionResource{}
	paths := trimPrefixSuffixSpace(strings.Split(u.Path, "/"))
	if len(paths) == 0 {
		return nil, "", "", fmt.Errorf("parse url error")
	}
	switch paths[0] {
	case "api":
		gvr.Group = ""
		gvr.Version = paths[1]
		gvr.Resource = paths[2]
		// /api/v1/watch/namespaces/{namespaces}/pods
		if len(paths) >= 6 && paths[3] == "namespaces" {
			namespace = paths[4]
		}
	case "apis":
		if paths[1] == "crd" {
			// /apis/crd/nuwa.nip.io/v1/waters?watch=1&resourceVersion=5738162
			remove(paths, 1)
		}
		gvr.Group = paths[1]
		gvr.Version = paths[2]
		gvr.Resource = paths[3]
		// /apis/apps/v1/watch/namespaces/{namespaces}/deployments
		if len(paths) >= 7 && paths[4] == "namespaces" {
			namespace = paths[5]
		}

	default:
		return nil, "", "", fmt.Errorf("parser url (%s) error", apiUrl)
	}

	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, "", "", err
	}
	if len(m["resourceVersion"]) > 0 {
		resourceVersion = m["resourceVersion"][0]
	}

	return
}

func trimPrefixSuffixSpace(slice []string) []string {
	for i, s := range slice {
		if s == "" {
			slice = slice[i+1:]
		}
	}
	return slice
}

func listenByApis(event *workloadservice.Event, g *gin.Context, eventChan chan Event) (closedSet []chan struct{}, err error) {
	apis := g.QueryArray("api")
	for _, api := range apis {
		gvr, ns, rv, err := parseForApiUrl(api)
		if err != nil {
			return closedSet, err
		}
		closed := make(chan struct{})
		k8sWatchChan, err := event.Watch(*gvr, ns, rv, 60, nil, closed)
		if err != nil {
			log.Printf("watch for gvr: %s stream error: %s for api request %s \r\n", gvr, err, api)
			continue
		}
		go func(ce chan Event) {
			for item := range k8sWatchChan {
				eventChan <- Event{Type: item.Type, Object: item.Object}
			}
		}(eventChan)
		closedSet = append(closedSet, closed)
	}
	return closedSet, nil
}

// watchStream watch api request resource group and the version
// after server timeout then close send closed event to client side server watcher close
func (w *WorkloadsAPI) WatchStream(g *gin.Context) {
	eventChan := make(chan Event, 32)
	closedSet, err := listenByApis(w.event, g, eventChan)
	shutdown := func() {
		for _, closed := range closedSet {
			closed <- struct{}{}
		}
	}
	if err != nil {
		shutdown()
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter"},
		)
		return
	}

	defer shutdown()

	g.Stream(func(w io.Writer) bool {
		select {
		case event, ok := <-eventChan:
			if !ok {
				g.SSEvent(
					"STREAM_END",
					&Event{Url: g.Request.URL.String(), Status: 410},
				)
				return false
			}
			g.SSEvent("", event)
		}
		return true
	})
}
