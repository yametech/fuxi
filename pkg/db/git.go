package db

import (
	"github.com/jinzhu/gorm"
	ops "github.com/yametech/fuxi/proto/ops"
)

type Git struct {
	//gorm.Model
	ops.Git
}

//func AutoMigrateGit() {
//	DB.AutoMigrate(&Git{})
//}

//FindGitByUserId find a git detail  by userid
func FindGitByUserId(uid int32) (*Git, error) {
	var git Git
	err := DB.Find(&git).Where("userid =  ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &git, err
}

//CreateGit create a git
func CreateGit(git Git) error {
	return DB.Create(&git).Error
}
