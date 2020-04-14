package workload

// PersistentVolume the kubernetes native resource persistent volume
type PersistentVolume struct {
	WorkloadsResourceHandler
}

// NewPersistentVolume exported
func NewPersistentVolume() *PersistentVolume {
	return &PersistentVolume{&defaultImplWorkloadsResourceHandler{}}
}
