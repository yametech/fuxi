package db

import (
	"github.com/jinzhu/gorm"
)

type Permission struct {
	gorm.Model
	Name     string `json:"name"`  // [超级管理员],[运维管理,业务运维],[开发主管,业务管理,业务开发]
	Value    uint32 `json:"value"` // [777],[770,077],[077,037,007]
	IsDelete bool   `json:"is_delete"`
	Comment  string `json:"comment"`
}

func AutoMigratePermission() {
	DB.AutoMigrate(&Permission{})
}

// Convert int32 type to Chinese expression
func (p *Permission) PermissionTransfer() []map[string]interface{} {
	var transfer []map[string]interface{}
	ResourceConstList := MergeResourceConstList()

	hasPer := func(v uint32, bit uint32) bool {
		if v&(1<<bit) > 0 {
			return true
		}
		return false
	}

	for i := range ResourceConstList {
		// set Chinese alias， bit and permission value
		chiAlias, _ := ShardingResourceChineseAliasList.Get(ResourceConstList[i])
		bit, _ := ShardingResourceList.Get(ResourceConstList[i])
		is := hasPer(p.Value, bit)
		// append to transfer
		transfer = append(transfer, map[string]interface{}{
			"name": ResourceConstList[i],
			"chi":  chiAlias,
			"is":   is,
		})
	}

	return transfer
}

// Convert  Chinese expression to int32 type
func (p *Permission) PermissionAuthorizeValue(config []interface{}) uint32 {

	var r uint32 = 0
	for i := range config {
		s, exists := ShardingResourceList.Get(config[i].(map[string]interface{})["name"].(string))
		if exists {
			var v uint32 = 0
			v = v | (1 << s)
			r += v
		}
	}

	return r
}
