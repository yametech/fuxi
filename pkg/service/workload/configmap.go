package workload

// ConfigMaps the kubernetes native resource configmaps
type ConfigMaps struct {
	WorkloadsResourceHandler
}

// NewConfigMaps exported
func NewConfigMaps() *ConfigMaps {
	return &ConfigMaps{&defaultImplWorkloadsResourceHandler{}}
}
