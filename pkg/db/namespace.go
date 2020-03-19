package db

import (
	"github.com/jinzhu/gorm"
	ns "github.com/yametech/fuxi/proto/ns"
)

type Namespace struct {
	gorm.Model
	ns.NS
}

func AutoMigrateNamespace() {
	DB.AutoMigrate(&Namespace{})
}

type excuteFunc func(name string) error

func CreateNamespace(ns *Namespace, excuteFunc excuteFunc) error {
	tx := DB.Begin()
	err := tx.Create(&ns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := excuteFunc(ns.Namespacename); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func NamespaceList() ([]Namespace, error) {
	var namespaces []Namespace
	err := DB.Find(&namespaces).Where("isdelete = ?", false).Error
	if err != nil {
		return nil, err
	}
	return namespaces, nil
}

func DeleteNamespace(nameSpaceName string, excuteFunc excuteFunc) error {
	var ns Namespace

	err := DB.Find(&ns).Error
	if err != nil {
		return err
	}
	ns.Isdelete = true

	tx := DB.Begin()
	err = tx.Update(ns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := excuteFunc(ns.Namespacename); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func EditNamespace(ns Namespace) error {
	err := DB.Update(ns).Error
	if err != nil {
		return err
	}
	return nil
}

func FindNamespaceByName(name string) (*Namespace, error) {
	ns := &Namespace{}
	err := DB.Find(ns).Where("namespacename = ?", name).Error
	if err != nil {
		return nil, err
	}
	return ns, nil
}
