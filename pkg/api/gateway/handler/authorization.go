package handler

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/yametech/fuxi/pkg/service/workload"

	"github.com/go-acme/lego/log"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/service/base"
	watch "k8s.io/apimachinery/pkg/watch"
)

type OpType string

const (
	POST   OpType = "POST"
	GET    OpType = "GET"
	PUT    OpType = "PUT"
	DELELE OpType = "DELELE"
)

type Resource struct {
	Op   OpType // eg: http restful[POST,GET,PUT,DELELE]
	Path string // eg: /workload/apis/nuwa.nip.io/v1/stones &&  /workload/apis/nuwa.nip.io/v1/namespaces/:namespace/stones/:name
}

type Role struct {
	Name      string   `json:"name"`
	Namespace []string `json:"namespace"`
	PermValue uint32   `json:"permValue"`
	baseDept  *base.BaseDepartment
}

func NewRole(name string, permValue uint32) *Role {
	role := Role{
		Name:      name,
		PermValue: permValue,
		baseDept:  base.NewBaseDepartment(),
	}

	//TODO 关联关系??
	return &role
}

type Roles []*Role

func (r Roles) search(roleName string) *Role {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Name <= r[j].Name
	})
	idx := sort.Search(len(r), func(i int) bool {
		return r[i].Name >= roleName
	})
	return r[idx]
}

type Authorization struct {
	baseRole *base.BaseRole
	mutex    sync.RWMutex
	password string
	roles    Roles `json:"roles"`
}

func NewAuthorization() (*Authorization, error) {
	auth := &Authorization{
		baseRole: base.NewBaseRole(),
		mutex:    sync.RWMutex{},
		roles:    make([]*Role, 0),
	}
	if err := auth.watch(); err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *Authorization) watch() error {
	list, err := a.baseRole.List("", "", 0, 0, nil)
	if err != nil {
		return err
	}
	roleList := &v1.BaseRoleList{}
	marshalData, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshalData, roleList)
	if err != nil {
		return err
	}

	for _, baseRole := range roleList.Items {
		a.roles = append(a.roles, &Role{
			Name: baseRole.Name,
		})
	}
	stream, err := a.baseRole.Watch("", roleList.GetResourceVersion(), 0, nil)
	if err != nil {
		return err
	}
	go func() {
		for event := range stream {
			obj := event.Object.(*v1.BaseRole)
			switch event.Type {
			case watch.Added:
				a.append(NewRole(obj.GetName(), obj.Spec.Value))
			case watch.Deleted:
				a.remove(obj.GetName())
			case watch.Modified:
				a.update(obj.GetName(), obj.Spec.Value)
			}
		}
	}()

	return nil
}

func (auth *Authorization) update(roleName string, permValue uint32) {
	auth.mutex.RLocker()
	defer auth.mutex.RUnlock()
	role := auth.roles.search(roleName)
	role.PermValue = permValue
	return
}

func (auth *Authorization) append(role *Role) {
	auth.mutex.RLocker()
	defer auth.mutex.RUnlock()
	auth.roles = append(auth.roles, role)
}

func (auth *Authorization) remove(roleName string) {
	auth.mutex.RLocker()
	defer auth.mutex.RUnlock()
	for index, _role := range auth.roles {
		if _role.Name == roleName {
			auth.roles = append(
				auth.roles[:index],
				auth.roles[index+1:]...,
			)
		}
	}
}

type AuthorizationStorage struct {
	mutex     sync.RWMutex
	data      map[string]*Authorization
	baseUser  *base.BaseUser
	namespace *workload.Namespace
}

func NewAuthorizationStorage() (*AuthorizationStorage, error) {
	authorizationStorage := &AuthorizationStorage{
		mutex:     sync.RWMutex{},
		data:      make(map[string]*Authorization),
		baseUser:  base.NewBaseUser(),
		namespace: workload.NewNamespace(),
	}
	return authorizationStorage, nil
}

// TODO 查找用户可以访问的空间
func (a *Authorization) UserAllowNamespace(user string) []string {
	if user == "admin" {

	}
	return []string{}
}

func (a *AuthorizationStorage) Auth(username string, password string) (bool, error) {
	auth := a.get(username)
	if auth == nil {
		return false, fmt.Errorf("%s", "user does not exist")
	}
	if auth.password != password {
		return false, fmt.Errorf("%s", "user password does not match")
	}
	return true, nil
}

func (a *AuthorizationStorage) watchUserData() error {
	userList, err := a.baseUser.List("", "", 0, 0, nil)
	if err != nil {
		return err
	}
	baseUserList := &v1.BaseUserList{}
	marshalData, err := json.Marshal(userList)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshalData, baseUserList)
	if err != nil {
		return err
	}

	for _, user := range userList.Items {
		authorization, err := NewAuthorization()
		if err != nil {
			return err
		}
		a.data[user.GetName()] = authorization
	}
	stream, err := a.baseUser.Watch("", userList.GetResourceVersion(), 0, nil)
	if err != nil {
		return err
	}
	go func() {
		for userEvent := range stream {
			obj := userEvent.Object.(*v1.BaseUser)
			switch userEvent.Type {
			case watch.Added:
				authorization, err := NewAuthorization()
				if err != nil {
					log.Infof("new authorization error: %s", err)
				}
				a.add(obj.GetName(), authorization)
			case watch.Deleted:
				a.remove(obj.GetName())
			}
		}
	}()
	return nil
}

func (a *AuthorizationStorage) get(name string) *Authorization {
	a.mutex.RLocker()
	defer a.mutex.RUnlock()
	author, ok := a.data[name]
	if !ok {
		return nil
	}
	return author
}

func (a *AuthorizationStorage) exist(user string) bool {
	a.mutex.RLocker()
	defer a.mutex.RUnlock()
	if _, ok := (a.data)[user]; !ok {
		return false
	}
	return true
}

func (a *AuthorizationStorage) add(user string, auth *Authorization) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.data[user] = auth
}

func (a *AuthorizationStorage) remove(user string) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	delete(a.data, user)
}
