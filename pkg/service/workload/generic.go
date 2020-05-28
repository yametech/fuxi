package workload

import "github.com/yametech/fuxi/pkg/service/common"

type Generic struct {
	common.WorkloadsResourceHandler
}

func NewGeneric() *Generic {
	return &Generic{&common.DefaultImplWorkloadsResourceHandler{}}
}
