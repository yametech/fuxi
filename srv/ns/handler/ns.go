package handler

import (
	"context"

	"github.com/yametech/fuxi/pkg/db"
	kubeclient "github.com/yametech/fuxi/pkg/kubernetes/clientv1"
	ns "github.com/yametech/fuxi/proto/ns"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Ns struct{}

//{
//"monitors":"10.200.100.200:6789,10.200.100.201:6789,10.200.100.202:6789,10.200.100.203:6789,10.200.100.204:6789",
//"adminSecretName":"ceph-secret",
//"pool":"kube",
//"userId":"kube",
//"userSecretName":"ceph-secret",
//"namespace":"",
//"adminId":"admin",
//"isdefault":true,
//"cidrblock" :"10.16.0.0/16",
//"namespacesexcludeips":"10.16.0.1,",
//"namespaces":"test,",
//"namespacename":"test",
//"cpu":"1",
//"memory":"200Mi",
//"storage":"1Gi",
//"cputhreshold":1,
//"memorythreshold":10,
//"storagethreshold":2
// "creator":"dba"
//}

func (ns *Ns) CreateNameSpace(ctx context.Context, req *ns.NS, rsp *ns.NSResponse) error {
	namespace := &db.Namespace{
		NS: *req,
	}
	if err := db.CreateNamespace(namespace, CreateNamespaceByKubeClient); err != nil {
		return err
	}

	return nil
}

func CreateNamespaceByKubeClient(name string) error {
	_, err := kubeclient.K8sClient.CoreV1().Namespaces().Create(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: name},
	})
	if err != nil {
		return err
	}
	return nil
}

func (ns *Ns) NamespaceList(ctx context.Context, in *ns.Empty, rsp *ns.NamespaceListResponse) error {
	namespaces, err := db.NamespaceList()
	if err != nil {
		return err
	}

	for _, namespace := range namespaces {
		rsp.Namespaces = append(rsp.Namespaces, &namespace.NS)
	}
	return nil
}

func DeleteNamespaceByKubeClient(name string) error {
	err := kubeclient.K8sClient.CoreV1().Namespaces().Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (ns *Ns) DeleteNamespace(ctx context.Context, req *ns.NamespaceName, rsp *ns.Empty) error {
	err := db.DeleteNamespace(req.Namespacename, DeleteNamespaceByKubeClient)
	if err != nil {
		return err
	}
	return nil
}

//todo:mabe update namespace name will update namespace name by kubeclient
func (ns *Ns) EditNamespace(ctx context.Context, req *ns.NS, rsp *ns.Empty) error {
	n := db.Namespace{NS: *req}
	err := db.EditNamespace(n)
	if err != nil {
		return err
	}
	return nil
}
