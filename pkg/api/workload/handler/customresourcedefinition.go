package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ListCustomResourceDefinition List CustomResourceDefinition
func (w *WorkloadsAPI) ListCustomResourceDefinition(g *gin.Context) {
	list, err := w.customResourceDefinition.List(dyn.ResourceCustomResourceDefinition, "", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	customResourceDefinitionList := &v1beta1.CustomResourceDefinitionList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, customResourceDefinitionList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, customResourceDefinitionList)
}

func (w *WorkloadsAPI) ListCustomResourceRouter() ([]string, error) {
	list, err := w.customResourceDefinition.List(dyn.ResourceCustomResourceDefinition, "", "", 0, 0, nil)
	if err != nil {
		return nil, err
	}

	customResourceDefinitionList := &v1beta1.CustomResourceDefinitionList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(marshalData, customResourceDefinitionList)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)
	for _, item := range customResourceDefinitionList.Items {
		group := fmt.Sprintf("%s/:version/%s", item.Spec.Group, item.Spec.Names.Plural)
		results = append(results, group)
	}
	return results, nil
}

// ListGeneralCustomResourceDefinition List General CustomResourceDefinition
func (w *WorkloadsAPI) ListGeneralCustomResourceDefinition(g *gin.Context) {
	u, err := url.Parse(g.Request.RequestURI)
	if err != nil {
		toRequestParamsError(g, err)
		return
	}
	paths := trimPrefixSuffixSpace(strings.Split(u.Path, "/"))
	if len(paths) < 5 {
		toRequestParamsError(g, fmt.Errorf("request url error %s", g.Request.RequestURI))
		return
	}
	group := paths[2]
	version := paths[3]
	resource := paths[4]

	//  import general GroupVersionResource
	groupVersionResource := schema.GroupVersionResource{Group: group, Version: version, Resource: resource}
	list, err := w.customResourceDefinition.List(groupVersionResource, "", "", 0, 0, nil)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	customResourceDefinitionList := &v1beta1.CustomResourceDefinitionList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, customResourceDefinitionList)
	if err != nil {
		toInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
