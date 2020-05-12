package workload

import (
	dyn "github.com/yametech/fuxi/pkg/kubernetes/client"
)

// ConfigMaps the kubernetes native resource configmaps
type ConfigMaps struct {
	WorkloadsResourceHandler
}

// NewConfigMaps exported
func NewConfigMaps() *ConfigMaps {
	return &ConfigMaps{&defaultImplWorkloadsResourceHandler{dyn.ResourceConfigMaps}}
}
