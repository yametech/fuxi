package db

import "github.com/jinzhu/gorm"

type RoleGroup struct {
	gorm.Model
	GroupName string
	RoleId    int
	UserId    int
	Status    int
	IsDelete  bool
	Creator   int
}

func AutoMigrateRoleGroup() {
	DB.AutoMigrate(&RoleGroup{})
}
