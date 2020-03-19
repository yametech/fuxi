package db

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model

	UserID      int32
	RoleGroupID int32
	OwnerUserID int32
	ParentID    int32
	DeptLevel   int32
	DeptType    int32
	Creator     int
	IsDelete    bool
	Deptlevel   int32

	DeptName    string `gorm:"type:varchar(100);unique_index;not null"`
	DeptFullID  string `gorm:"size:1024"`
	Description string `gorm:"size:1024"`
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

//FindDepartment find a Department
func FindDepartment(uid int32) (Department, error) {
	var d Department
	err := DB.Find(&d).Where("UserID = ?", uid).Error
	if err != nil {
		return d, err
	}
	return d, err
}

func DeleteDepartment(d Department) error {
	err := DB.Delete(&d).Error
	if err != nil {
		return err
	}
	return nil
}

func EditDepartment(d Department) error {
	err := DB.Update(&d).Error
	if err != nil {
		return err
	}
	return nil
}

func DepartmentList() ([]*Department, error) {
	var ds []*Department
	err := DB.Find(&ds).Error
	if err != nil {
		return nil, err
	}
	return ds, nil
}
