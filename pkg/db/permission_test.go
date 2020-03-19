package db

import (
	"fmt"
	"testing"
)

func TestShardingResourceList(t *testing.T) {

	u1, b1 := ShardingResourceList.Get("APP_VIEW")
	fmt.Println(u1, b1)
	if b1 && u1 != APP_VIEW_BIT {
		t.Fatal("non my expect")
	}

	u2, b2 := ShardingResourceList.Get("APP_UPDATE")
	fmt.Println(u2, b2)
	if b2 && u2 != APP_UPDATE_BIT {
		t.Fatal("non my expect")
	}
}

func TestMergeResourceList(t *testing.T) {

	mergeList := MergeResourceConstList()
	for i := range mergeList {
		fmt.Println(mergeList[i])
		fmt.Println(ShardingResourceList.Get(mergeList[i]))
		fmt.Println(ShardingResourceChineseAliasList.Get(mergeList[i]))
	}

}

func TestPermission_PermissionTransfer(t *testing.T) {

	//p := Permission{Value: 196608}
	p := Permission{Value: 327694}
	fmt.Println(p.PermissionTransfer())
}

func TestPermission_PermissionAuthorize(t *testing.T) {

	p := Permission{}
	config := make([]interface{}, 0)
	//
	config = append(config, map[string]interface{}{"name": "NS_VIEW"})
	config = append(config, map[string]interface{}{"name": "NS_UPDATE"})
	config = append(config, map[string]interface{}{"name": "APP_UPDATE"})
	config = append(config, map[string]interface{}{"name": "APP_CREATE"})
	config = append(config, map[string]interface{}{"name": "APP_DELETE"})

	//
	fmt.Println(config)
	fmt.Println(p.PermissionAuthorizeValue(config))

}
