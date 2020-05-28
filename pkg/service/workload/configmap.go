package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ConfigMaps the kubernetes native resource configmaps
type ConfigMaps struct {
	common.WorkloadsResourceHandler
}

// NewConfigMaps exported
func NewConfigMaps() *ConfigMaps {
	return &ConfigMaps{&common.DefaultImplWorkloadsResourceHandler{dyn.ResourceConfigMaps}}
}
