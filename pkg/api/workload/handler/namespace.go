package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	nuwav1 "github.com/yametech/nuwa/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

type patchAnnotateData struct {
	Namespace string   `json:"namespace"`
	Nodes     []string `json:"nodes"`
}

func jsonPatchNodeListData(o interface{}) (string, error) {
	bs, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (w *WorkloadsAPI) PatchAnnotateNodeNamespace(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	pad := patchAnnotateData{}
	err = json.Unmarshal(rawData, &pad)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	cords := make(nuwav1.Coordinates, 0)
	for _, nodeName := range pad.Nodes {
		obj, err := w.node.Get("", nodeName)
		if err != nil {
			common.ToRequestParamsError(g, err)
			return
		}
		nodeUnstructured := obj.(*unstructured.Unstructured)
		node := &corev1.Node{}
		if err = runtime.DefaultUnstructuredConverter.FromUnstructured(nodeUnstructured.Object, node); err != nil {
			common.ToRequestParamsError(g, err)
			return
		}
		cord := nuwav1.Coordinate{}
		if labels := node.GetLabels(); labels != nil {
			zone, exist := labels[nuwav1.NuwaZoneFlag]
			if !exist {
				common.ToInternalServerError(g, "", fmt.Errorf("node %s not label nuwa zone", node.GetName()))
				return
			}
			cord.Zone = zone
			rack, exist := labels[nuwav1.NuwaRackFlag]
			if !exist {
				common.ToInternalServerError(g, "", fmt.Errorf("node %s not label nuwa rack", node.GetName()))
				return
			}
			cord.Rack = rack
			host, exist := labels[nuwav1.NuwaHostFlag]
			if !exist {
				common.ToInternalServerError(g, "", fmt.Errorf("node %s not label nuwa host", node.GetName()))
				return
			}
			cord.Host = host
		}
		cords = append(cords, cord)
	}
	patchNodeListValue, err := jsonPatchNodeListData(cords)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	patchData := map[string]interface{}{
		"metadata": map[string]interface{}{
			"annotations": map[string]string{
				nuwav1.NuwaLimitFlag: patchNodeListValue,
			},
		},
	}
	_, err = w.namespace.Patch("", pad.Namespace, patchData)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, "")
}

// Create Namespace
func (w *WorkloadsAPI) CreateNamespace(g *gin.Context) {
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := corev1.Namespace{}
	err = json.Unmarshal(rawData, &obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}
	newObj, err := w.namespace.Apply("", obj.Name, unstructuredStruct)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}

// Delete Namespace
func (w *WorkloadsAPI) DeleteNamespace(g *gin.Context) {
	namespaceName := g.Param("namespace")
	err := w.namespace.Delete("", namespaceName)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, "")
}

// Get Namespace
func (w *WorkloadsAPI) GetNamespace(g *gin.Context) {
	namespaceName := g.Param("namespace")
	item, err := w.namespace.Get("", namespaceName)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Namespaces
func (w *WorkloadsAPI) ListNamespace(g *gin.Context) {
	list, err := w.namespace.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	namespaceList := &corev1.NamespaceList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	err = json.Unmarshal(marshalData, namespaceList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, namespaceList)
}
