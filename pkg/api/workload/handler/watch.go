package handler

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
)

type WorkloadsEvent struct {
	Type   watch.EventType `json:"type"`
	Object runtime.Object  `json:"object"`
}

func (w *WorkloadsAPI) Watch(g *gin.Context) {

}
