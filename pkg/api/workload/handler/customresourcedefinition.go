package handler

import (
	"encoding/json"
	"fmt"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ListCustomResourceDefinition List CustomResourceDefinition
func (w *WorkloadsAPI) ListCustomResourceDefinition(g *gin.Context) {
	//selector := metav1.LabelSelector{MatchLabels: map[string]string{"name": "statefulsets.nuwa.nip.io"}}
	list, err := w.customResourceDefinition.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	customResourceDefinitionList := &v1beta1.CustomResourceDefinitionList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, customResourceDefinitionList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, customResourceDefinitionList)
}

func (w *WorkloadsAPI) ListCustomResourceRouter(gvrString []string) ([]string, error) {
	list, err := w.customResourceDefinition.List("", "", 0, 0, nil)
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
		var apiVersionUrl string
		for _, version := range item.Spec.Versions {
			apiVersionUrl = fmt.Sprintf("%s/%s/%s", item.Spec.Group, version.Name, item.Spec.Names.Plural)
			needIgnore := false
			for _, ignoreItem := range gvrString {
				if apiVersionUrl == ignoreItem {
					needIgnore = true
				}
			}
			if !needIgnore {
				results = append(results, apiVersionUrl)
			}
		}
	}
	return results, nil
}

// ListGeneralCustomResourceDefinition List General CustomResourceDefinition
func (w *WorkloadsAPI) ListGeneralCustomResourceDefinition(g *gin.Context) {
	u, err := url.Parse(g.Request.RequestURI)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	paths := trimPrefixSuffixSpace(strings.Split(u.Path, "/"))
	if len(paths) < 5 {
		common.ToRequestParamsError(g, fmt.Errorf("request url error %s", g.Request.RequestURI))
		return
	}
	group := paths[2]
	version := paths[3]
	resource := paths[4]

	//  import general GetGroupVersionResource
	groupVersionResource := schema.GroupVersionResource{Group: group, Version: version, Resource: resource}
	w.generic.SetGroupVersionResource(groupVersionResource)
	list, err := w.generic.List("", "", 0, 0, nil)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	customResourceDefinitionList := &v1beta1.CustomResourceDefinitionList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, customResourceDefinitionList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, list)
}
