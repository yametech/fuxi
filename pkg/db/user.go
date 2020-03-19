package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name          *string   `binding:"exists" form:"name" json:"name" example:"test" swaggertype:"string" gorm:"name;type:varchar(64);unique_index;not null"`
	Password      *string   `binding:"exists" form:"password" json:"password" example:"test" swaggertype:"string"  gorm:"type:varchar(64);not null"`
	RoleId        int       `form:"role_id" json:"role_id"`
	DepartmentId  int       `form:"department_id" json:"department_id"`
	LastLoginTime time.Time `form:"last_login_time" json:"last_login_time" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatorId     int       `form:"creator_id" json:"creator_id"`
	IsDelete      bool      `form:"is_delete" json:"is_delete"`
	Email         string    `form:"email" json:"email" gorm:"type:varchar(64); unique_index"`
	Display       string    `form:"display" json:"display" gorm:"type:varchar(200)"`
	IsAdmin       bool      `form:"is_admin" json:"is_admin"`
}

func (User) TableName() string { return "user" }
func AutoMigrateUser()         { DB.AutoMigrate(&User{}) }

//FindUserIdByName find user id by Name
func FindUserIdByName(name string) (int32, error) {
	var user User
	err := DB.Find(&user).Where("name = ?", name).Error
	if err != nil {
		return 0, err
	}
	return int32(user.ID), err
}
