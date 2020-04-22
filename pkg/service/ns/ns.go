package ns

import (
	informers "github.com/alauda/kube-ovn/pkg/client/informers/externalversions"
	ovnclient "github.com/yametech/fuxi/pkg/ovn"
	"time"
)

type NS interface {
	Subnet
}
type NSService struct {
	informer informers.SharedInformerFactory
}

func NewNS(defaultResync time.Duration) *NSService {
	return &NSService{
		informer: informers.NewSharedInformerFactory(ovnclient.KubeOvnClient, defaultResync),
	}
}

func (ns *NSService) Start(stopCh <-chan struct{}) {
	ns.informer.Start(stopCh)
	ns.informer.WaitForCacheSync(stopCh)
}

var _ NS = (*NSService)(nil)
