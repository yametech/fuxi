package db

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	RoleName        string
	RolePermGroupid int
	IsDelete        bool
	Creator         int
}

func AutoMigrateRole() {
	DB.AutoMigrate(&Role{})
}
