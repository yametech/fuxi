package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yametech/fuxi/pkg/api/common"
	"net/http"
	"strconv"
)

const (
	// operator
	VIEW   = 1 << iota // 0000 0001  = 1
	UPDATE             // 0000 0010  = 2
	CREATE             // 0000 0100  = 4
	DELETE             // 0000 1000  = 8

	//     8bit  |
	// 0000 0000 | 0000 0111 | 0000 0111 | 0111 0111     = uint32
	// byte=uint8

	// operator app/workload bit
	APP_VIEW_BIT   = 0
	APP_UPDATE_BIT = 1
	APP_CREATE_BIT = 2
	APP_DELETE_BIT = 3

	APP_VIEW   = "APP_VIEW"
	APP_UPDATE = "APP_UPDATE"
	APP_CREATE = "APP_CREATE"
	APP_DELETE = "APP_DELETE"

	// user default
	USER_VIEW_BIT   = 4
	USER_UPDATE_BIT = 5
	USER_CREATE_BIT = 6
	USER_DELETE_BIT = 7

	USER_VIEW   = "USER_VIEW"
	USER_UPDATE = "USER_UPDATE"
	USER_CREATE = "USER_CREATE"
	USER_DELETE = "USER_DELETE"

	// operator service bit
	SRV_VIEW_BIT   = 8
	SRV_UPDATE_BIT = 9
	SRV_CREATE_BIT = 10
	SRV_DELETE_BIT = 11

	SRV_VIEW   = "SRV_VIEW"
	SRV_UPDATE = "SRV_UPDATE"
	SRV_CREATE = "SRV_CREATE"
	SRV_DELETE = "SRV_DELETE"

	// operator permission bit
	PERMISSION_VIEW_BIT   = 12
	PERMISSION_UPDATE_BIT = 13
	PERMISSION_CREATE_BIT = 14
	PERMISSION_DELETE_BIT = 15

	PERMISSION_VIEW   = "PERMISSION_VIEW"
	PERMISSION_UPDATE = "PERMISSION_UPDATE"
	PERMISSION_CREATE = "PERMISSION_CREATE"
	PERMISSION_DELETE = "PERMISSION_DELETE"

	// operator namespace bit
	NS_VIEW_BIT   = 16
	NS_UPDATE_BIT = 17
	NS_CREATE_BIT = 18
	NS_DELETE_BIT = 19

	NS_VIEW   = "NS_VIEW"
	NS_UPDATE = "NS_UPDATE"
	NS_CREATE = "NS_CREATE"
	NS_DELETE = "NS_DELETE"

	// operator service bit
	// RS_VIEW_BIT         = 24
	// RS_PERM_UPDATE_BIT  = 25
	// RS_PERM_Create__BIT = 26
	// RS_PERM_DELETE_BIT  = 27
)

var ShardingResourceList *ResourceList

type keyValuePair struct {
	k string
	v int
}

type ResourceList struct {
	v map[string]uint32
}

type KeyValuePairChi struct {
	k string
	v string
}

type ResourceChiList struct {
	v map[string]string
}

func init() {

	// inti global ShardingResourceList var
	if ShardingResourceList == nil {
		ShardingResourceList = &ResourceList{v: make(map[string]uint32)}
	}
	ShardingResourceList.push(
		// APP
		keyValuePair{APP_VIEW, APP_VIEW_BIT},
		keyValuePair{APP_UPDATE, APP_UPDATE_BIT},
		keyValuePair{APP_CREATE, APP_CREATE_BIT},
		keyValuePair{APP_DELETE, APP_DELETE_BIT},
		// USER
		keyValuePair{USER_VIEW, USER_VIEW_BIT},
		keyValuePair{USER_UPDATE, USER_UPDATE_BIT},
		keyValuePair{USER_CREATE, USER_CREATE_BIT},
		keyValuePair{USER_DELETE, USER_DELETE_BIT},
		// SRV
		keyValuePair{SRV_VIEW, SRV_VIEW_BIT},
		keyValuePair{SRV_UPDATE, SRV_UPDATE_BIT},
		keyValuePair{SRV_CREATE, SRV_CREATE_BIT},
		keyValuePair{SRV_DELETE, SRV_DELETE_BIT},
		// PERMISSION
		keyValuePair{PERMISSION_VIEW, PERMISSION_VIEW_BIT},
		keyValuePair{PERMISSION_UPDATE, PERMISSION_CREATE_BIT},
		keyValuePair{PERMISSION_CREATE, PERMISSION_CREATE_BIT},
		keyValuePair{PERMISSION_DELETE, PERMISSION_DELETE_BIT},
		// NS
		keyValuePair{NS_VIEW, NS_VIEW_BIT},
		keyValuePair{NS_UPDATE, NS_CREATE_BIT},
		keyValuePair{NS_CREATE, NS_CREATE_BIT},
		keyValuePair{NS_DELETE, NS_DELETE_BIT},
	)
}

func MergeResourceConstList() []string {
	TypeList := [...]string{"NS", "SRV", "APP", "PERMISSION"}
	OperateList := [...]string{"_VIEW", "_UPDATE", "_CREATE", "_DELETE"}
	var this []string

	for i := range TypeList {
		for j := range OperateList {
			// append merge string
			this = append(this, TypeList[i]+OperateList[j])
		}
	}

	return this
}

func (r *ResourceList) Get(s string) (uint32, bool) {
	value, exist := r.v[s]
	if !exist {
		return 0, false
	}
	return value, true
}

func (r *ResourceList) push(args ...keyValuePair) {
	for i := range args {
		r.v[args[i].k] = uint32(args[i].v)
	}
}

func (r *ResourceChiList) Get(s string) (string, bool) {
	value, exist := r.v[s]
	if !exist {
		return "", false
	}
	return value, true
}

func (r *ResourceChiList) push(args ...KeyValuePairChi) {
	for i := range args {
		r.v[args[i].k] = args[i].v
	}
}

func CheckPermission(user uint32, has ...func(uint32) bool) bool {
	for i := range has {
		f := has[i]
		if !f(user) {
			return false
		}
	}
	return true
}

// Convert  Chinese expression to int32 type
func PermissionAuthorizeValue(config []interface{}) uint32 {

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

func (b *BaseAPI) PermissionTransfer(g *gin.Context) {

	p, err := strconv.Atoi(g.Param("value"))
	if err != nil {
		common.ToRequestParamsError(g, err)
		return
	}

	permValue := uint32(p)
	var transfer []string
	ResourceConstList := MergeResourceConstList()

	hasPer := func(v uint32, bit uint32) bool {
		if v&(1<<bit) > 0 {
			return true
		}
		return false
	}

	for i := range ResourceConstList {
		bit, _ := ShardingResourceList.Get(ResourceConstList[i])
		is := hasPer(permValue, bit)
		// append to transfer
		if is {
			transfer = append(transfer, ResourceConstList[i])
		}
	}

	g.JSON(http.StatusOK, transfer)
}
