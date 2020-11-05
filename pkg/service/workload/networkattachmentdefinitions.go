package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// NetworkAttachmentDefinition the kubernetes native service account
type NetworkAttachmentDefinition struct {
	common.WorkloadsResourceHandler
}

// NewNetworkAttachmentDefinition exported
func NewNetworkAttachmentDefinition() *NetworkAttachmentDefinition {
	return &NetworkAttachmentDefinition{&common.DefaultImplWorkloadsResourceHandler{
		types.ResourceNetworkAttachmentDefinition,
	}}
}
