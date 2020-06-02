package workload

import (
	"github.com/yametech/fuxi/pkg/kubernetes/types"
	"github.com/yametech/fuxi/pkg/service/common"
)

// ConfigMaps the kubernetes native resource configmaps
type ConfigMaps struct {
	common.WorkloadsResourceHandler
}

// NewConfigMaps exported
func NewConfigMaps() *ConfigMaps {
	return &ConfigMaps{&common.DefaultImplWorkloadsResourceHandler{types.ResourceConfigMaps}}
}
