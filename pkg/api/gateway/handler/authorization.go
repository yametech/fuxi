package handler

import (
	"encoding/json"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/service/base"
	watch "k8s.io/apimachinery/pkg/watch"
	"sort"
	"sync"
)

type Role struct {
	Name      string   `json:"name"`
	Namespace []string `json:"namespace"`
	PermValue int32    `json:"permValue"`
}
type Roles []Role

func (r Roles) search(roleName string) Role {
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
	roles    Roles `json:"roles"`
}

func NewAuthorization() *Authorization {
	auth := &Authorization{
		baseRole: base.NewBaseRole(),
		mutex:    sync.RWMutex{},
		roles:    make([]Role, 0),
	}
	return auth
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
		a.roles = append(a.roles, Role{
			Name: baseRole.Name,
		})
	}
	stream, err := a.baseRole.Watch("", roleList.GetResourceVersion(), 0, nil)
	if err != nil {
		return err
	}
	go func() {
		for event := range stream {
			switch event {
			//Added    EventType = "ADDED"
			//Modified EventType = "MODIFIED"
			//Deleted  EventType = "DELETED"
			//Bookmark EventType = "BOOKMARK"
			//Error    EventType = "ERROR"
			}
		}
	}()
	return nil
}

func (auth *Authorization) update(roleName string, namespaces []string, permValue int32) {
	auth.mutex.RLocker()
	defer auth.mutex.RUnlock()
	role := auth.roles.search(roleName)
	role.Namespace = namespaces
	role.PermValue = permValue
	return
}

func (auth *Authorization) append(role Role) {
	auth.mutex.RLocker()
	defer auth.mutex.RUnlock()
	auth.roles = append(auth.roles, role)
}

func (auth *Authorization) remove(role Role) {
	auth.mutex.RLocker()
	defer auth.mutex.RUnlock()
	for index, _role := range auth.roles {
		if _role.Name == role.Name {
			auth.roles = append(
				auth.roles[:index],
				auth.roles[index+1:]...,
			)
		}
	}
}

type AuthorizationStorage struct {
	mutex sync.RWMutex
	data  map[string]*Authorization
	user  *base.BaseUser
	dept  *base.BaseDepartment
}

func NewAuthorizationStorage() (*AuthorizationStorage, error) {
	authorizationStorage := &AuthorizationStorage{
		mutex: sync.RWMutex{},
		data:  make(map[string]*Authorization),
		user:  base.NewBaseUser(),
	}
	return authorizationStorage, nil
}

func (a *AuthorizationStorage) watchUserData() error {
	userList, err := a.user.List("", "", 0, 0, nil)
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
		a.data[user.GetName()] = NewAuthorization()
	}
	stream, err := a.user.Watch("", userList.GetResourceVersion(), 0, nil)
	if err != nil {
		return err
	}
	go func() {
		for userEvent := range stream {
			obj := userEvent.Object.(*v1.BaseUser)
			switch userEvent.Type {
			case watch.Added:
				a.add(obj.GetName(), NewAuthorization())
			case watch.Deleted:
				a.remove(obj.GetName())
			}
		}
	}()
	return nil
}

func (a *AuthorizationStorage) Exist(user string) bool {
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
