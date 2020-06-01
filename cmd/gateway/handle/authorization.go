package handle

type User string

type Role struct {
	Namespace      string `json:"namespace"`
	PermissionData int64  `json:"permission_data"`
}

type Authorization struct {
	Roles []Role `json:"roles"`
}

type AuthorizationStorage map[User]Authorization

func (a *AuthorizationStorage) Exist(user string) bool {
	if _, ok := (*a)[(User)(user)]; !ok {
		return false
	}
	return true
}
