package handler

import (
	"context"
	"net/http"

	"github.com/yametech/fuxi/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	ns "github.com/yametech/fuxi/proto/ns"
)

//NSService
type NSService struct {
	nsC ns.NsService
}

//New new  OpsService
func New(client client.Client) *NSService {
	return &NSService{nsC: ns.NewNsService("go.micro.srv.ns", client)}
}

func (n *NSService) CreateNamespace(c *gin.Context) {
	log.Info("NSService NS.CreateNamespace API request")
	ns := db.Namespace{}
	if err := c.ShouldBindJSON(&ns); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	_, err := n.nsC.CreateNameSpace(context.TODO(), &ns.NS)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "titile": "CreateNamespace"})
	}
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

func (n *NSService) NamespaceList(c *gin.Context) {
	log.Info("NSService NS.NamespaceList API request")
	resp, err := n.nsC.NamespaceList(context.TODO(), &ns.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "titile": "CreateNamespace"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp.Namespaces})
}

func (n *NSService) DeleteNamespace(c *gin.Context) {
	log.Info("NSService NS.CreateNamespace API request")
	name := c.Param("namespacename")
	_, err := n.nsC.DeleteNamespace(context.TODO(), &ns.NamespaceName{Namespacename: name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "titile": "DeleteNamespace"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (n *NSService) EditNamespace(c *gin.Context) {
	log.Info("NSService NS.CreateNamespace API request")
	ns := db.Namespace{}
	if err := c.ShouldBindJSON(&ns); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	_, err := n.nsC.EditNamespace(context.TODO(), &ns.NS)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "titile": "CreateNamespace"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
