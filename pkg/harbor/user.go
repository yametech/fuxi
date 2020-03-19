package harbor

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

const (
	GETUSER = "/api/users/"
)

var (
	CreateUserError    = errors.New("create a user error")
	DeleteUserError    = errors.New("delete a user error")
	UpdateUserPassword = errors.New("update a user password error")
)

type User struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Realname     string `json:"realname"`
	Comment      string `json:"comment"`
	Deleted      bool   `json:"deleted"`
	RoleName     string `json:"role_name"`
	RoleID       int    `json:"role_id"`
	HasAdminRole bool   `json:"has_admin_role"`
	ResetUUID    string `json:"reset_uuid"`
	Salt         string `json:"Salt"`
	CreationTime string `json:"creation_time"`
	UpdateTime   string `json:"update_time"`
}

type UsersSearchFliter struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type UpdateUser struct {
	UserId      string `json:"user_id"`
	OldPassWord string `json:"old_password"`
	NewPassWord string `json:"new_password"`
}

//GetUsers Get registered users of Harbor.
func (c *HarborClient) GetUsers(userSearchFliter *UsersSearchFliter) ([]User, error) {
	var users []User
	res, err := c.Client.R().SetPathParams(
		map[string]string{
			"username":  userSearchFliter.UserName,
			"email":     userSearchFliter.Email,
			"page":      strconv.Itoa(userSearchFliter.Page),
			"page_size": strconv.Itoa(userSearchFliter.PageSize),
		}).Get(GETUSER)
	if err != nil {
		return users, err
	}
	err = json.Unmarshal(res.Body(), &users)
	if err != nil {
		return users, err
	}
	return users, nil
}

//CreateUser   Creates a new user account.
func (c *HarborClient) CreateUser(user *User) error {
	res, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		Post(GETUSER)
	if err != nil {
		return err
	}
	if res.StatusCode() == 201 {
		return err
	}
	return errors.WithStack(CreateUserError)
}

//DeleteUser  Mark a registered user as be removed.
func (c *HarborClient) DeleteUser(userId int) error {
	uId := strconv.Itoa(userId)
	res, err := c.Client.R().Delete(GETUSER + uId)
	if err != nil {
		return errors.WithStack(err)
	}
	if res.StatusCode() == 200 {
		return nil
	}
	return errors.WithStack(DeleteUserError)
}

//UpdateUserPassword update a user password
func (c *HarborClient) UpdateUserPassword(updateUser *UpdateUser) error {
	res, err := c.Client.R().
		SetBody(updateUser).
		Put(GETUSER + updateUser.UserId + "/password")
	if err != nil {
		return err
	}
	if res.StatusCode() == 200 {
		return nil
	}
	return errors.Wrap(UpdateUserPassword, string(res.Body()))
}
