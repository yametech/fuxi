package workload

// StorageClass the kubernetes native resource storage class
type StorageClass struct {
	WorkloadsResourceHandler
}

// NewStorageClass exported
func NewStorageClass() *StorageClass {
	return &StorageClass{&defaultImplWorkloadsResourceHandler{}}
}
