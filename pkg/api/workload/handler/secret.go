package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	serviceCommon "github.com/yametech/fuxi/pkg/service/common"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

// Get Secret
func (w *WorkloadsAPI) GetSecret(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.secret.Get(namespace, name)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List Secret
func (w *WorkloadsAPI) ListSecret(g *gin.Context) {
	var list *unstructured.UnstructuredList
	var err error

	namespace := g.Param("namespace")
	if namespace == "" {
		list, err = w.secret.List("", "", 0, 0, nil)
	} else {
		labelSelector := fmt.Sprintf("hide!=%s", "1")
		list, err = w.secret.List(namespace, "", 0, 0, labelSelector)
	}
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	secretList := &v1.SecretList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, secretList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, secretList)
}

// List Ops Secret
func (w *WorkloadsAPI) ListOpsSecret(g *gin.Context) {
	var list *unstructured.UnstructuredList
	var err error

	namespace := g.Param("namespace")
	labelSelector := fmt.Sprintf("tektonConfig=%s", "1")
	if namespace == "" {
		list, err = w.secret.List("", "", 0, 0, labelSelector)
	} else {
		list, err = w.secret.List(namespace, "", 0, 0, labelSelector)
	}
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	secretList := &v1.SecretList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, secretList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, secretList)
}

// Create Secret
func (w *WorkloadsAPI) CreateSecret(g *gin.Context) {
	namespace := g.Param("namespace")
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	obj := v1.Secret{}
	err = json.Unmarshal(rawData, &obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	if obj.Type == v1.SecretTypeDockerConfigJson {
		config := make(map[string]map[string]string)
		err := json.Unmarshal(obj.Data[".dockerconfigjson"], &config)
		if err != nil {
			common.ToInternalServerError(g, "", err)
			return
		}
		for address, args := range config {

			bytesData, err := serviceCommon.HandleDockerCfgJSONContent(
				args["username"], args["password"], args["email"], address)

			if err != nil {
				common.ToInternalServerError(g, "", err)
				return
			}
			obj.Data = map[string][]byte{".dockerconfigjson": bytesData}
			delete(obj.Labels, ".dockerconfigjson")
			delete(obj.Annotations, ".dockerconfigjson")
		}
	}

	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	unstructuredStruct := &unstructured.Unstructured{
		Object: unstructuredObj,
	}

	newObj, err := w.secret.Apply(namespace, obj.Name, unstructuredStruct)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}

	g.JSON(http.StatusOK, newObj)
}
