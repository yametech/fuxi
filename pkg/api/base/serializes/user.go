package serializes

import "github.com/yametech/fuxi/pkg/db"

type User db.User

func (u *User) UserRoleSerializer() JSON {
	return JSON{
		"id": u.ID,
	}
}
