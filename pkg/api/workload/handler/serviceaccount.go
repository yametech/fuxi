package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	v1 "k8s.io/api/core/v1"
	"net/http"
)

type patchServiceAccountSecret struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func removeObjectReference(slice []v1.ObjectReference, name string) []v1.ObjectReference {
	tmpMap := make(map[string]v1.ObjectReference)
	for i := range slice {
		tmpMap[slice[i].Name] = slice[i]
	}
	delete(tmpMap, name)
	result := make([]v1.ObjectReference, 0)
	for _, v := range tmpMap {
		result = append(result, v)
	}
	return result
}

func (w *WorkloadsAPI) PatchSecretServiceAccount(g *gin.Context) {
	method := g.Param("method")
	rawData, err := g.GetRawData()
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	pad := patchServiceAccountSecret{}
	err = json.Unmarshal(rawData, &pad)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	serviceAccountUnstructed, err := w.serviceAccount.Get(pad.Namespace, "default")
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	serviceAccount := &v1.ServiceAccount{}
	if err := common.RuntimeObjectToInstanceObj(serviceAccountUnstructed, serviceAccount); err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	if method == "add" {
		serviceAccount.Secrets = append(serviceAccount.Secrets, v1.ObjectReference{Name: pad.Name})
	} else {
		serviceAccount.Secrets = removeObjectReference(serviceAccount.Secrets, pad.Name)
	}

	unstructured, err := common.InstanceToUnstructured(serviceAccount)
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}
	_, _, err = w.serviceAccount.Apply(pad.Namespace, "default", unstructured)
	if err != nil {
		common.ToInternalServerError(g, serviceAccount, err)
		return
	}
	g.JSON(http.StatusOK, nil)
}

// Get ServiceAccount
func (w *WorkloadsAPI) GetServiceAccount(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.serviceAccount.Get(namespace, name)
	if err != nil {
		common.ResourceNotFoundError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, item)
}

// List ServiceAccount
func (w *WorkloadsAPI) ListServiceAccount(g *gin.Context) {
	list, err := resourceList(g, w.serviceAccount)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	serviceAccountList := &v1.ServiceAccountList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	err = json.Unmarshal(marshalData, serviceAccountList)
	if err != nil {
		common.ToInternalServerError(g, "", err)
		return
	}
	g.JSON(http.StatusOK, serviceAccountList)
}
