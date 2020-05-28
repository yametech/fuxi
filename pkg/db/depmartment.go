package db

import (
	"github.com/jinzhu/gorm"
)

type Department struct {
	gorm.Model

	UserID      int32  `form:"user_id" json:"user_id" gorm:"user_id;type:int"`
	RoleGroupID int32  `form:"role_group_id" json:"role_group_id" gorm:"role_group_id;type:int"`
	OwnerUserID int32  `form:"owner_user_id" json:"owner_user_id" gorm:"owner_user_id;type:int"`
	ParentID    int32  `form:"parent_id" json:"parent_id" gorm:"parent_id;type:int"`
	DeptLevel   int32  `form:"dept_level" json:"dept_level" gorm:"dept_level;type:int"`
	DeptType    int32  `form:"dept_type" json:"dept_type" gorm:"dept_type;type:int"`
	Creator     int    `form:"creator" json:"creator" gorm:"creator;type:int"`
	IsDelete    bool   `form:"is_delete" json:"is_delete" gorm:"is_delete;type:bool"`
	DeptName    string `form:"dept_name" json:"dept_name" gorm:"type:varchar(100);unique_index;not null"`
	DeptFullID  string `form:"dept_full_id" json:"dept_full_id" gorm:"size:1024"`
	Description string `form:"description" json:"description" gorm:"size:1024"`
}

type DepartmentSerializer struct {
	id       int32
	userName string
	name     string
}

func (d Department) TableName() string {
	return "department"
}

func AutoMigrateDepartment() {
	DB.AutoMigrate(&Department{})
}

func CreateDepartment(Department Department) error {
	err := DB.Create(Department).Error
	if err != nil {
		return err
	}
	return err
}

// Find
func FindDepartment(uid int32) (Department, error) {
	var d Department
	err := DB.Find(&d).Where("UserID = ?", uid).Error
	if err != nil {
		return d, err
	}
	return d, err
}

// Get
func GetUserNameByUserId(userId int32) string {
	var u User
	err := DB.First(&u, userId).Error
	if err != nil {
		return ""
	}
	return *u.Name
}

// Delete
func DeleteDepartment(d Department) error {
	err := DB.Delete(&d).Error
	if err != nil {
		return err
	}
	return nil
}

// Edit
func EditDepartment(d Department) error {
	err := DB.Update(&d).Error
	if err != nil {
		return err
	}
	return nil
}

// List
func DepartmentList() ([]*Department, error) {
	var ds []*Department
	err := DB.Find(&ds).Error
	if err != nil {
		return nil, err
	}
	return ds, nil
}
