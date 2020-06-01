package handler

import (
	"github.com/yametech/fuxi/pkg/service/base"
)

// BaseAPI all resource operate
type BaseAPI struct {
	basedepartments *base.BaseDepartment
	basepermissions *base.BasePermission
	baseroles       *base.BaseRole
	baseusers       *base.BaseUser
	baseroleusers   *base.BaseRoleUser
	baseroleperms   *base.BaseRolePerm
}

func NewBaseAPi() *BaseAPI {
	return &BaseAPI{
		basedepartments: base.NewBaseDepartment(),
		basepermissions: base.NewBasePermission(),
		baseroles:       base.NewBaseRole(),
		baseusers:       base.NewBaseUser(),
		baseroleusers:   base.NewBaseRoleUser(),
		baseroleperms:   base.NewBaseRolePerm(),
	}
}
