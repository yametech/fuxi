package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"net/http"
)

func (w *WorkloadsAPI) GetDeployment(g *gin.Context) {
	namespace := g.Param("namespace")
	name := g.Param("name")
	item, err := w.deployments.Get(dyn.ResourceDeployment, namespace, name)
	if err != nil {
		panic(err)
	}
	g.JSON(http.StatusOK, item)
}

func (w *WorkloadsAPI) ListDeployment(g *gin.Context) {
	list, _ := w.deployments.List(dyn.ResourceDeployment, "", "", 0, 100, nil)
	deploymentList := &appsv1.DeploymentList{}
	data, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(data, deploymentList)
	g.JSON(http.StatusOK, deploymentList)
}

func (w *WorkloadsAPI) ApplyDeployment(g *gin.Context) {
	var formData map[string]interface{}
	if err := g.BindJSON(&formData); err != nil {
		g.JSON(http.StatusBadRequest,
			gin.H{
				code:   http.StatusBadRequest,
				data:   "",
				msg:    err.Error(),
				status: "Request bad parameter",
			},
		)
		return
	}
	unstructuredData := &unstructured.Unstructured{Object: formData}
	md, _ := formData["metadata"]
	metadata := md.(map[string]interface{})
	namespace := metadata["namespace"].(string)
	name := metadata["name"].(string)
	err := w.deployments.Apply(dyn.ResourceDeployment, namespace, name, unstructuredData)
	if err != nil {
		g.JSON(http.StatusInternalServerError,
			gin.H{
				code:   http.StatusInternalServerError,
				data:   "",
				msg:    err.Error(),
				status: "apply error",
			},
		)
		return
	}

	ds := []unstructured.Unstructured{
		*unstructuredData,
	}

	g.JSON(http.StatusOK, ds)
}

func (w *WorkloadsAPI) DeleteDeployment(g *gin.Context) {}
