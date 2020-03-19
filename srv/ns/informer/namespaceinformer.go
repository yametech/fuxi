package informer

import (
	"fmt"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"
	"github.com/yametech/fuxi/pkg/db"
	kubeclient "github.com/yametech/fuxi/pkg/k8s/client"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"

	ovnv1 "github.com/alauda/kube-ovn/pkg/apis/kubeovn/v1"
	ovnclient "github.com/yametech/fuxi/pkg/ovn"
	storagev1 "k8s.io/api/storage/v1"
)

type NamespaceInformer struct {
	informerFactory   informers.SharedInformerFactory
	namespaceInformer coreinformers.NamespaceInformer
}

//when namespace create will create quota ,according to department
func (c *NamespaceInformer) nsAdd(obj interface{}) {
	ns := obj.(*corev1.Namespace)
	glog.Infof("ns CREATED: %s", ns.Name)
	fmt.Fprintf(os.Stdout, "create ns %s\n", ns.Name)
	if ns.Name != "test" {
		return
	}
	n, err := db.FindNamespaceByName(ns.Name)
	if err != nil {
		glog.Error("FindNamespaceByName error: ", err)
	}
	if err := createCephStoreClass(*n); err != nil {
		glog.Error("createCephStoreClass error: ", err)
	}
	if err := createOvnCidr(*n); err != nil {
		glog.Error("createOvnCidr error: ", err)
	}
	if err := createQutaResource(*n); err != nil {
		glog.Error("createQutaResource error: ", err)
	}
}

func (c *NamespaceInformer) nsUpdate(old, new interface{}) {
	//oldNs := old.(*v1.Namespace)
	//newNs := new.(*v1.Namespace)
	glog.Infof("ns update")
}

func (c *NamespaceInformer) nsDelete(obj interface{}) {
	ns := obj.(*corev1.Namespace)
	glog.Infof("ns DELETED: %s/%s", ns.Name)
}

func (c *NamespaceInformer) Run(stopCh chan struct{}) error {
	// Starts all the shared informers that have been created by the factory so
	// far.
	c.informerFactory.Start(stopCh)
	// wait for the initial synchronization of the local cache.
	if !cache.WaitForCacheSync(stopCh, c.namespaceInformer.Informer().HasSynced) {
		return fmt.Errorf("Failed to sync")
	}
	return nil
}

// NewNamespaceLoggingController creates a nsLoggingController
func NewNamespaceLoggingController(informerFactory informers.SharedInformerFactory) *NamespaceInformer {
	nsInformer := informerFactory.Core().V1().Namespaces()

	c := &NamespaceInformer{
		informerFactory:   informerFactory,
		namespaceInformer: nsInformer,
	}
	nsInformer.Informer().AddEventHandler(
		// Your custom resource event handlers.
		cache.ResourceEventHandlerFuncs{
			// Called on creation
			AddFunc: c.nsAdd,
			// Called on resource update and every resyncPeriod on existing resources.
			UpdateFunc: c.nsUpdate,
			// Called on resource deletion.
			DeleteFunc: c.nsDelete,
		},
	)
	return c
}

//todo:use hard-coded,which will config it
func createCephStoreClass(ns db.Namespace) error {

	annotations := make(map[string]string)
	params := make(map[string]string)
	if ns.GetIsdefault() {
		annotations["storageclass.kubernetes.io/is-default-class"] = "true"
		//todo:jude superadmin
		//params["adminId"] = ns.GetAdminId()
		//params["adminSecretName"] = ns.GetAdminSecretName()
		//params["adminSecretNamespace"] = "default"
	}

	params["monitors"] = ns.GetMonitors()

	params["pool"] = ns.GetPool()
	params["userId"] = ns.GetUserId()
	params["userSecretName"] = ns.GetUserSecretName()
	params["imageFormat"] = "2"
	params["imageFeatures"] = "layering"

	_, err := kubeclient.K8sClient.StorageV1().StorageClasses().Create(&storagev1.StorageClass{
		TypeMeta:    v1.TypeMeta{},
		ObjectMeta:  v1.ObjectMeta{Annotations: annotations, Namespace: ns.Namespacename, Name: ns.Creator},
		Provisioner: "ceph.com/rbd",
		Parameters:  params,
	})
	if err != nil {
		return err
	}
	return nil
}

func createOvnCidr(ns db.Namespace) error {
	_, err := ovnclient.KubeOvnClient.KubeovnV1().Subnets().Create(&ovnv1.Subnet{
		ObjectMeta: v1.ObjectMeta{Namespace: ns.Namespacename, Name: ns.Creator},
		Spec: ovnv1.SubnetSpec{
			Default:    false,
			Protocol:   "IPv4",
			Namespaces: strings.Split(ns.Namespaces, ","),
			CIDRBlock:  ns.Cidrblock,
			ExcludeIps: strings.Split(ns.Namespacesexcludeips, ","),
		},
		Status: ovnv1.SubnetStatus{},
	})
	if err != nil {
		return err
	}
	return nil
}

func createQutaResource(ns db.Namespace) error {
	resourceList := corev1.ResourceList{}
	resourceList[corev1.ResourceCPU] = resource.MustParse(ns.Cpu)
	resourceList[corev1.ResourceMemory] = resource.MustParse(ns.Memory)
	//resourceList[corev1.ResourceStorage] = resource.MustParse(ns.Storage)
	_, err := kubeclient.K8sClient.CoreV1().ResourceQuotas(ns.Namespacename).Create(&corev1.ResourceQuota{
		ObjectMeta: v1.ObjectMeta{Namespace: ns.Namespacename, Name: ns.Creator},
		Spec: corev1.ResourceQuotaSpec{
			Hard: resourceList,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
