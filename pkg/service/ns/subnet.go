package ns

import (
	ovnv1 "github.com/alauda/kube-ovn/pkg/apis/kubeovn/v1"
	ovnclient "github.com/yametech/fuxi/pkg/ovn"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type Subnet interface{
	CreateSubnet(sn SubNet) error
	SubNetDelete(name string) error
	SubNetUpdate(sn SubNet) error
    SubNetList() ([]*SubNet,error)
	GetSubNet(name string)(*SubNet,error)
}


type SubNet struct {
	//Namespace string
	Name string
	IsDefault bool
	Protocol string
	Namespaces []string
	CIDRBlock string
	ExcludeIps []string
	Gateway string
	GatewayType string
	GatewayNode string
	NatOutgoing bool
	Private bool
	AllowSubnets [] string
}

//CreateSubnet one cluster only one default subnet if Default set to true
func (ns *NSService)CreateSubnet(sn SubNet) error {

	_, err := ovnclient.KubeOvnClient.KubeovnV1().Subnets().Create(&ovnv1.Subnet{
		ObjectMeta: v1.ObjectMeta{Name: sn.Name},
		Spec: ovnv1.SubnetSpec{
			Default:      sn.IsDefault,
			Protocol:     sn.Protocol,
			Namespaces:   sn.Namespaces,
			CIDRBlock:    sn.CIDRBlock,
			Gateway:      sn.Gateway,
			ExcludeIps:   sn.ExcludeIps,
			GatewayType:  sn.GatewayType,
			GatewayNode:  sn.GatewayNode,
			NatOutgoing:  sn.NatOutgoing,
			Private:      sn.Private,
			AllowSubnets: sn.AllowSubnets,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

//SubNetDelete delete a subnet
func  (ns *NSService)SubNetDelete(name string) error {

	err := ovnclient.KubeOvnClient.KubeovnV1().Subnets().Delete(name,&metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

//SubNetUpdate update a subnet config
func  (ns *NSService)SubNetUpdate(sn SubNet) error {

	_, err :=ovnclient.KubeOvnClient.KubeovnV1().Subnets().Update(
		&ovnv1.Subnet{
			ObjectMeta: v1.ObjectMeta{Name: sn.Name},
			Spec: ovnv1.SubnetSpec{
				Default:      sn.IsDefault,
				Protocol:     sn.Protocol,
				Namespaces:   sn.Namespaces,
				CIDRBlock:    sn.CIDRBlock,
				Gateway:      sn.Gateway,
				ExcludeIps:   sn.ExcludeIps,
				GatewayType:  sn.GatewayType,
				GatewayNode:  sn.GatewayNode,
				NatOutgoing:  sn.NatOutgoing,
				Private:      sn.Private,
				AllowSubnets: sn.AllowSubnets,
			},
		})
	if err != nil {
		return err
	}
	return nil
}

//SubNetList select all subnets
func  (ns *NSService)SubNetList() ([]*SubNet,error) {

	ret,err:=ns.informer.Kubeovn().V1().Subnets().Lister().List(labels.NewSelector())
	if err != nil {
		return nil, err
	}
	var sbs []*SubNet
	for _, subnet := range ret {
		sbs = append(sbs,&SubNet{
			Name:        subnet.Name,
			IsDefault:    subnet.Spec.Default,
			Protocol:     subnet.Spec.Protocol,
			Namespaces:   subnet.Spec.Namespaces,
			CIDRBlock:    subnet.Spec.CIDRBlock,
			ExcludeIps:   subnet.Spec.ExcludeIps,
			Gateway:      subnet.Spec.Gateway,
			GatewayType:  subnet.Spec.GatewayType,
			GatewayNode:  subnet.Spec.GatewayNode,
			NatOutgoing:  subnet.Spec.NatOutgoing,
			Private:      subnet.Spec.Private,
			AllowSubnets: subnet.Spec.AllowSubnets,
		})
	}
	return sbs, nil
}

//GetSubNet get subnet
func  (ns *NSService)GetSubNet(name string)(*SubNet,error){
	subnet, err :=ovnclient.KubeOvnClient.KubeovnV1().Subnets().Get(name,metav1.GetOptions{})
	if err != nil {
		return nil,err
	}

	return &SubNet{
		Name:        subnet.Name,
		IsDefault:    subnet.Spec.Default,
		Protocol:     subnet.Spec.Protocol,
		Namespaces:   subnet.Spec.Namespaces,
		CIDRBlock:    subnet.Spec.CIDRBlock,
		ExcludeIps:   subnet.Spec.ExcludeIps,
		Gateway:      subnet.Spec.Gateway,
		GatewayType:  subnet.Spec.GatewayType,
		GatewayNode:  subnet.Spec.GatewayNode,
		NatOutgoing:  subnet.Spec.NatOutgoing,
		Private:      subnet.Spec.Private,
		AllowSubnets: subnet.Spec.AllowSubnets,
	},nil
}



