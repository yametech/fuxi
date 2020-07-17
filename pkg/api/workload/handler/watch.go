package handler

import (
	"fmt"
	constraint_common "github.com/yametech/fuxi/common"
	"github.com/yametech/fuxi/pkg/service/common"
	"io"
	corev1 "k8s.io/api/core/v1"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	workloadservice "github.com/yametech/fuxi/pkg/service/workload"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
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
	case "api": //  eg: /api/v1/pods
		gvr.Group = ""
		gvr.Version = paths[1]
		gvr.Resource = paths[2]
		if len(paths) == 5 {
			//  eg: /api/v1/namespaces/dxp/pods
			gvr.Resource = paths[4]
			namespace = paths[3]
		}
	case "apis": // eg: /apis/apps/v1/deployments
		gvr.Group = paths[1]
		gvr.Version = paths[2]
		gvr.Resource = paths[3]
		// eg: /apis/apps/v1/namespaces/dxp/deployments
		if len(paths) == 6 {
			gvr.Resource = paths[5]
			namespace = paths[4]
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

func listenByApis(event *workloadservice.Generic, g *gin.Context, eventChan chan Event, closed chan struct{}) {
	defer close(eventChan)
	apis := g.QueryArray("api")
	wg := sync.WaitGroup{}
	wg.Add(len(apis))
	for _, api := range apis {
		gvr, ns, rv, err := parseForApiUrl(api)
		if err != nil {
			log.Printf("parser api url error %s", api)
			return
		}
		event.SetGroupVersionResource(*gvr)
		if rv == "" {
			log.Printf("watch for gvr: %s stream error: %s for api request %s \r\n", gvr, err, api)
			continue
		}

		var k8sWatchChan <-chan watch.Event
		// Redirect fuxi.nip.io/workload resources
		if ns != "" && gvr.Group == "fuxi.nip.io" && gvr.Resource == "workloads" {
			k8sWatchChan, err = event.Watch(constraint_common.WorkloadsDeployTemplateNamespace, rv, 0, fmt.Sprintf("namespace=%s", ns))
		} else if gvr.Resource == "ops-secrets" {
			gvr.Resource = "secrets"
			event.SetGroupVersionResource(*gvr)
			k8sWatchChan, err = event.Watch(ns, rv, 0, fmt.Sprintf("tektonConfig=%s", "1"))
		} else {
			k8sWatchChan, err = event.Watch(ns, rv, 0, nil)
		}
		if err != nil {
			log.Printf("watch for gvr: %s stream error: %s for api request %s \r\n", gvr, err, api)
			continue
		}

		go func() {
			defer wg.Done()
			for {
				select {
				case _, ok := <-closed:
					if !ok {
						return
					}
				case item, ok := <-k8sWatchChan:
					if !ok {
						return
					}
					// ignore all error
					newObj, _ := insectionObject(item.Object)
					eventChan <- Event{
						Type:   item.Type,
						Object: newObj,
					}
				}
			}
		}()
	}
	wg.Wait()
}

// watchStream watch api request resource group and the version
// after server timeout then close send closed event to clientv2 side server watcher close
func (w *WorkloadsAPI) WatchStream(g *gin.Context) {
	eventChan := make(chan Event, 32)
	closed := make(chan struct{})
	go listenByApis(w.generic, g, eventChan, closed)

	streamEndEvent := Event{
		Type:   watch.EventType("STREAM_END"),
		Url:    g.Request.URL.String(),
		Status: 410,
	}
	g.Stream(func(w io.Writer) bool {
		select {
		case <-g.Writer.CloseNotify():
			close(closed)
			g.SSEvent("", streamEndEvent)
			return false
		case event, ok := <-eventChan:
			if !ok {
				g.SSEvent("", streamEndEvent)
				return false
			}
			g.SSEvent("", event)
		}
		return true
	})
}

func insectionObject(object runtime.Object) (runtime.Object, error) {
	if object.GetObjectKind().GroupVersionKind().Kind == "Secret" {
		secret := &corev1.Secret{}
		if err := common.RuntimeObjectToInstanceObj(object, secret); err != nil {
			return object, err
		}
		labels := secret.GetLabels()
		if _, exist := labels["tektonConfig"]; !exist {
			return object, nil
		}
		selfLink := secret.GetSelfLink()
		secret.SetSelfLink(strings.Replace(selfLink, "/secrets", "/ops-secrets", 1))
		return secret, nil
	}
	return object, nil
}
