package handler

import (
	"github.com/yametech/fuxi/pkg/service/base"
)

// BaseAPI all resource operate
type BaseAPI struct {
	basedepartments *base.BaseDepartment
	baseroles       *base.BaseRole
	baseusers       *base.BaseUser
	baseroleusers   *base.BaseRoleUser
}

func NewBaseAPi() *BaseAPI {
	return &BaseAPI{
		basedepartments: base.NewBaseDepartment(),
		baseroles:       base.NewBaseRole(),
		baseusers:       base.NewBaseUser(),
		baseroleusers:   base.NewBaseRoleUser(),
	}
}
