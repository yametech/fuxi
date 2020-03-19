package db

import "github.com/jinzhu/gorm"

type ResourceType = string

type Resource struct {
	gorm.Model
	Uid          int64
	Name         string
	resourceType ResourceType
}
