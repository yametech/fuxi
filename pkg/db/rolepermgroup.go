package db

import "github.com/jinzhu/gorm"

type RolePermissionGroup struct {
	gorm.Model
	name         string
	PermissionId int
	IsDelete     bool
	Status       int
}

func AutoMigrateRolePermissionGroup() {
	DB.AutoMigrate(&RolePermissionGroup{})
}

func CreateRolePermissionGroup(rolePermissionGroup RolePermissionGroup) error {
	err := DB.Create(rolePermissionGroup).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteRolePermissionGroup(rolePermissionGroup RolePermissionGroup) error {
	err := DB.Delete(rolePermissionGroup).Error
	if err != nil {
		return err
	}
	return nil
}

func FindRolePermissionGroupById(id int) (*RolePermissionGroup, error) {
	var rolePermissionGroup RolePermissionGroup
	err := DB.Find(&rolePermissionGroup).Where("id =  ?", id).Error
	if err != nil {
		return nil, err
	}
	return &rolePermissionGroup, err
}

func UpdateRolePermissionGroup(rolePermissionGroup RolePermissionGroup) error {
	err := DB.Update(&rolePermissionGroup).Error
	if err != nil {
		return err
	}
	return nil
}

func RolePermissionGroupList() ([]*RolePermissionGroup, error) {
	var rolePermissionGroups []*RolePermissionGroup
	err := DB.Find(&rolePermissionGroups).Error
	if err != nil {
		return nil, err
	}
	return rolePermissionGroups, err
}
