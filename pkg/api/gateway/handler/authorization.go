package handler

import (
	"encoding/json"
	"fmt"
	"github.com/yametech/fuxi/thirdparty/lib/token"
	"sort"
	"time"

	"github.com/yametech/fuxi/common"
	v1 "github.com/yametech/fuxi/pkg/apis/fuxi/v1"
	"github.com/yametech/fuxi/pkg/service/base"
	"github.com/yametech/fuxi/pkg/service/workload"
	"k8s.io/apimachinery/pkg/runtime"
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
	userServices      *base.BaseUser
	roleServices      *base.BaseRole
	deptServices      *base.BaseDepartment
	namespaceServices *workload.Namespace
}

func NewAuthorization() *Authorization {
	auth := &Authorization{
		userServices:      base.NewBaseUser(),
		roleServices:      base.NewBaseRole(),
		deptServices:      base.NewBaseDepartment(),
		namespaceServices: workload.NewNamespace(),
	}
	return auth
}

func (auth *Authorization) allowNamespaceAccess(userName string, namespace string) (bool, error) {
	obj, err := auth.userServices.Get(common.BaseServiceStoreageNamespace, userName)
	if err != nil {
		return false, err
	}
	baseUser := &v1.BaseUser{}
	if err = runtimeObjectToInstanceObj(obj, baseUser); err != nil {
		return false, err
	}

	deptObj, err := auth.deptServices.Get(common.BaseServiceStoreageNamespace, baseUser.Spec.DepartmentId)
	if err != nil {
		return false, err
	}
	baseDept := &v1.BaseDepartment{}
	err = runtimeObjectToInstanceObj(deptObj, baseDept)
	if err != nil {
		return false, err
	}
	// not contain in allowNamespace
	if !In(baseDept.Spec.Namespace, namespace) {
		return false, nil
	}

	return true, nil
}

func In(list []string, item string) bool {
	for _, _item := range list {
		if _item == item {
			return true
		}
	}
	return false
}

func (auth *Authorization) getUser(userName string) (*v1.BaseUser, error) {
	obj, err := auth.userServices.Get(common.BaseServiceStoreageNamespace, userName)
	if err != nil {
		return nil, err
	}
	baseUser := &v1.BaseUser{}
	err = runtimeObjectToInstanceObj(obj, baseUser)
	if err != nil {
		return nil, err
	}
	return baseUser, nil
}

func (auth *Authorization) getDept(deptName string) (*v1.BaseDepartment, error) {
	obj, err := auth.deptServices.Get(common.BaseServiceStoreageNamespace, deptName)
	if err != nil {
		return nil, err
	}
	baseDept := &v1.BaseDepartment{}
	err = runtimeObjectToInstanceObj(obj, baseDept)
	if err != nil {
		return nil, err
	}
	return baseDept, nil
}

func (auth *Authorization) Config(tokenStr string) ([]byte, error) {
	return nil, nil
}

func (auth *Authorization) Auth(username, password string) ([]byte, error) {
	baseUser, err := auth.getUser(username)
	if err != nil {
		return nil, err
	}
	if *baseUser.Spec.Password != password {
		return nil, fmt.Errorf("password not match")
	}

	expireTime := time.Now().Add(time.Hour * 24).Unix()
	tokenStr, err := (&token.Token{}).Encode(common.MicroSaltUserHeader, username, expireTime)
	if err != nil {
		return nil, err
	}
	// user AllowedNamespaces
	dept, err := auth.getDept(baseUser.Spec.DepartmentId)
	if err != nil {
		return nil, err
	}
	return []byte(
		newUserConfig(
			username,
			tokenStr,
			dept.Spec.Namespace,
			dept.Spec.DefaultNamespace,
		).String(),
	), nil
}

func runtimeObjectToInstanceObj(robj runtime.Object, targeObj interface{}) error {
	bytesData, err := json.Marshal(robj)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytesData, targeObj)
}
