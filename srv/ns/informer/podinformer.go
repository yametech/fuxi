package informer

import (
	"log"

	"github.com/golang/glog"
	"github.com/yametech/fuxi/pkg/db"

	"github.com/yametech/fuxi/pkg/k8s/client"
	kubeclient "github.com/yametech/fuxi/pkg/k8s/client"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	v1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type PodInformer struct {
	podsLister v1.PodLister
	podsSynced cache.InformerSynced
}

func NewPodInformer() *PodInformer {
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(client.K8sClient, 0,
		kubeinformers.WithTweakListOptions(func(listOption *metav1.ListOptions) {
			listOption.AllowWatchBookmarks = true
		}))

	podInformer := informerFactory.Core().V1().Pods()

	podIn := &PodInformer{
		podsLister: podInformer.Lister(),
		podsSynced: podInformer.Informer().HasSynced,
		//addPodQueue:    workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "AddPod"),
		//deletePodQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "DeletePod"),
		//updatePodQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "UpdatePod"),
	}

	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    podIn.AddPod,
		DeleteFunc: podIn.DeletePod,
		UpdateFunc: podIn.UpdatePod,
	})

	return podIn

}

//AddPod  when Pod Add,which is willing check Quota,
// and then send alert to email DingTalk.
func (c *PodInformer) AddPod(obj interface{}) {
	pod := obj.(*corev1.Pod)
	quotaResources, err := kubeclient.K8sClient.CoreV1().ResourceQuotas(pod.GetNamespace()).List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	ns, err := db.FindNamespaceByName(pod.GetNamespace())
	if err != nil {
		glog.Error("FindQuotaAlertByDepartmentName error: ", err)
	}
	//will go to database,will find
	if int(ns.Memorythreshold) == quotaResources.Items[0].Status.Used.Memory().Sign() {
		//will alert something
	}
	if int(ns.Cputhreshold) == quotaResources.Items[0].Status.Used.Cpu().Sign() {
		//will alert something
	}
	if int(ns.Storagethreshold) == quotaResources.Items[0].Status.Used.StorageEphemeral().Sign() {
		//will alert something
	}

}

//DeletePod
func (c *PodInformer) DeletePod(obj interface{}) {
}

//UpdatePod
func (c *PodInformer) UpdatePod(oldObj, newObj interface{}) {
}
