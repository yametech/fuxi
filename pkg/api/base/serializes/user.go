package serializes

import "github.com/yametech/fuxi/pkg/db"

<<<<<<< HEAD
func (u *db.User) UserRoleSerializer() JSON {
=======
type User db.User

func (u *User) UserRoleSerializer() JSON {
>>>>>>> 0c71acc6e0202644d124b914f3c302d8c1d93ea5
	return JSON{
		"id": u.ID,
	}
}
