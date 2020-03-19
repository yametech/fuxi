package db

import "fmt"

const (
	// operator
	VIEW   = 1 << iota // 0000 0001  = 1
	UPDATE             // 0000 0010  = 2
	CREATE             // 0000 0100  = 4
	DELETE             // 0000 1000  = 8

	//     8bit  |
	// 0000 0000 | 0000 0111 | 0000 0111 | 0111 0111     = uint32
	// byte=uint8

	// fix ResourceType and Resource conflict
	NS  ResourceType = "NameSpace"
	SRV ResourceType = "Service"
	APP ResourceType = "APP"

	// operator app/workload bit
	APP_VIEW_BIT   = 0
	APP_UPDATE_BIT = 1
	APP_CREATE_BIT = 2
	APP_DELETE_BIT = 3

	APP_VIEW   = "APP_VIEW"
	APP_UPDATE = "APP_UPDATE"
	APP_CREATE = "APP_CREATE"
	APP_DELETE = "APP_DELETE"

	APP_VIEW_CHI   = "APP视图"
	APP_UPDATE_CHI = "APP更新"
	APP_CREATE_CHI = "APP创建"
	APP_DELETE_CHI = "APP删除"

	// user default
	USER_VIEW_BIT   = 4
	USER_UPDATE_BIT = 5
	USER_CREATE_BIT = 6
	USER_DELETE_BIT = 7

	USER_VIEW   = "USER_VIEW"
	USER_UPDATE = "USER_UPDATE"
	USER_CREATE = "USER_CREATE"
	USER_DELETE = "USER_DELETE"

	USER_VIEW_CHI   = "用户视图"
	USER_UPDATE_CHI = "用户更新"
	USER_CREATE_CHI = "用户创建"
	USER_DELETE_CHI = "用户删除"

	// operator service bit
	SRV_VIEW_BIT   = 8
	SRV_UPDATE_BIT = 9
	SRV_CREATE_BIT = 10
	SRV_DELETE_BIT = 11

	SRV_VIEW   = "SRV_VIEW"
	SRV_UPDATE = "SRV_UPDATE"
	SRV_CREATE = "SRV_CREATE"
	SRV_DELETE = "SRV_DELETE"

	SRV_VIEW_CHI   = "服务视图"
	SRV_UPDATE_CHI = "服务更新"
	SRV_CREATE_CHI = "服务创建"
	SRV_DELETE_CHI = "服务删除"

	// operator permission bit
	PERMISSION_VIEW_BIT   = 12
	PERMISSION_UPDATE_BIT = 13
	PERMISSION_CREATE_BIT = 14
	PERMISSION_DELETE_BIT = 15

	PERMISSION_VIEW   = "PERMISSION_VIEW"
	PERMISSION_UPDATE = "PERMISSION_UPDATE"
	PERMISSION_CREATE = "PERMISSION_CREATE"
	PERMISSION_DELETE = "PERMISSION_DELETE"

	PERMISSION_VIEW_CHI   = "权限视图"
	PERMISSION_UPDATE_CHI = "权限更新"
	PERMISSION_CREATE_CHI = "权限创建"
	PERMISSION_DELETE_CHI = "权限删除"

	// operator namespace bit
	NS_VIEW_BIT   = 16
	NS_UPDATE_BIT = 17
	NS_CREATE_BIT = 18
	NS_DELETE_BIT = 19

	NS_VIEW   = "NS_VIEW"
	NS_UPDATE = "NS_UPDATE"
	NS_CREATE = "NS_CREATE"
	NS_DELETE = "NS_DELETE"

	NS_VIEW_CHI   = "命名空间视图"
	NS_UPDATE_CHI = "命名空间更新"
	NS_CREATE_CHI = "命名空间创建"
	NS_DELETE_CHI = "命名空间删除"

	// operator service bit
	// RS_VIEW_BIT         = 24
	// RS_PERM_UPDATE_BIT  = 25
	// RS_PERM_Create__BIT = 26
	// RS_PERM_DELETE_BIT  = 27
)

var ShardingResourceList *ResourceList
var ShardingResourceChineseAliasList *ResourceChiList

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
	// init global ShardingResourceChineseAlias var
	if ShardingResourceChineseAliasList == nil {
		ShardingResourceChineseAliasList = &ResourceChiList{v: make(map[string]string)}
	}
	ShardingResourceChineseAliasList.push(
		// APP
		KeyValuePairChi{APP_VIEW, APP_VIEW_CHI},
		KeyValuePairChi{APP_UPDATE, APP_UPDATE_CHI},
		KeyValuePairChi{APP_CREATE, APP_CREATE_CHI},
		KeyValuePairChi{APP_DELETE, APP_DELETE_CHI},
		// USER
		KeyValuePairChi{USER_VIEW, USER_VIEW_CHI},
		KeyValuePairChi{USER_UPDATE, USER_UPDATE_CHI},
		KeyValuePairChi{USER_CREATE, USER_CREATE_CHI},
		KeyValuePairChi{USER_DELETE, USER_DELETE_CHI},
		// SRV
		KeyValuePairChi{SRV_VIEW, SRV_VIEW_CHI},
		KeyValuePairChi{SRV_UPDATE, SRV_CREATE_CHI},
		KeyValuePairChi{SRV_CREATE, SRV_CREATE_CHI},
		KeyValuePairChi{SRV_DELETE, SRV_DELETE_CHI},
		// PERMISSION
		KeyValuePairChi{PERMISSION_VIEW, PERMISSION_VIEW_CHI},
		KeyValuePairChi{PERMISSION_UPDATE, PERMISSION_UPDATE_CHI},
		KeyValuePairChi{PERMISSION_CREATE, PERMISSION_CREATE_CHI},
		KeyValuePairChi{PERMISSION_DELETE, PERMISSION_DELETE_CHI},
		// NS
		KeyValuePairChi{NS_VIEW, NS_VIEW_CHI},
		KeyValuePairChi{NS_UPDATE, NS_UPDATE_CHI},
		KeyValuePairChi{NS_CREATE, NS_CREATE_CHI},
		KeyValuePairChi{NS_DELETE, NS_DELETE_CHI},
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

func HasUserView(v uint32) bool {
	if v&(1<<USER_VIEW_BIT) > 0 {
		return true
	}
	return false
}
func HasUserUpdate(v uint32) bool {
	if v&(1<<USER_UPDATE_BIT) > 0 {
		return true
	}
	return false
}
func HasUserCreate(v uint32) bool {
	if v&(1<<USER_CREATE_BIT) > 0 {
		return true
	}
	return false
}
func HasUserDelete(v uint32) bool {
	if v&(1<<USER_DELETE_BIT) > 0 {
		return true
	}
	return false
}

// service
func HasSrvView(v uint32) bool {
	if v&(1<<SRV_VIEW_BIT) > 0 {
		return true
	}
	return false
}
func HasSrvUpdate(v uint32) bool {
	if v&(1<<SRV_UPDATE_BIT) > 0 {
		return true
	}
	return false
}
func HasSrvCreate(v uint32) bool {
	if v&(1<<SRV_CREATE_BIT) > 0 {
		return true
	}
	return false
}
func HasSrvDelete(v uint32) bool {
	if v&(1<<SRV_DELETE_BIT) > 0 {
		return true
	}
	return false
}

// permissions
func HasPermissionView(v uint32) bool {
	if v&(1<<PERMISSION_VIEW_BIT) > 0 {
		return true
	}
	return false
}
func HasPermissionUpdate(v uint32) bool {
	if v&(1<<PERMISSION_UPDATE_BIT) > 0 {
		return true
	}
	return false
}
func HasPermissionCreate(v uint32) bool {
	if v&(1<<PERMISSION_CREATE_BIT) > 0 {
		return true
	}
	return false
}
func HasPermissionDelete(v uint32) bool {
	if v&(1<<PERMISSION_DELETE_BIT) > 0 {
		return true
	}
	return false
}

// App
func HasAppView(v uint32) bool {
	if v&(1<<APP_VIEW_BIT) > 0 {
		return true
	}
	return false
}
func HasAppUpdate(v uint32) bool {
	if v&(1<<APP_UPDATE_BIT) > 0 {
		return true
	}
	return false
}
func HasAppCreate(v uint32) bool {
	if v&(1<<APP_CREATE_BIT) > 0 {
		return true
	}
	return false
}
func HasAppDelete(v uint32) bool {
	if v&(1<<APP_DELETE_BIT) > 0 {
		return true
	}
	return false
}

// Namespace
func HasNsView(v uint32) bool {
	if v&(1<<NS_VIEW_BIT) > 0 {
		return true
	}
	return false
}
func HasNsUpdate(v uint32) bool {
	if v&(1<<NS_UPDATE_BIT) > 0 {
		return true
	}
	return false
}
func HasNsCreate(v uint32) bool {
	if v&(1<<NS_CREATE_BIT) > 0 {
		return true
	}
	return false
}
func HasNsDelete(v uint32) bool {
	if v&(1<<NS_DELETE_BIT) > 0 {
		return true
	}
	return false
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

func test_example() {
	//  data := db.get_user_permission()	// 460551 = 070707
	// check the opeart user data with on current auth user
	if !CheckPermission(460551, HasUserView, HasNsUpdate) {
		fmt.Println("test success")
	}
}
