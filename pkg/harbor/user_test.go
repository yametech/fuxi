package harbor

import "testing"

func TestHarborClient_GetUsers(t *testing.T) {
	client := NewHarborClient("localhost/", "admin", "Harbor12345")

	users, err := client.GetUsers(&UsersSearchFliter{
		UserName: "",
		Email:    "",
		Page:     1,
		PageSize: 5,
	})
	if err != nil {
		t.Log(err)
	}
	t.Log(users)
}

func TestHarborClient_CreateUser(t *testing.T) {
	client := NewHarborClient("localhost/", "admin", "Harbor12345")

	err := client.CreateUser(&User{
		UserID:       8,
		Username:     "san.zhan8",
		Email:        "san.zhang8@163.com",
		Password:     "Harbor12345",
		Realname:     "san.zhang8",
		Comment:      "I%27m Zhang San",
		Deleted:      false,
		RoleName:     "",
		RoleID:       2,
		HasAdminRole: false,
		ResetUUID:    "",
		Salt:         "",
		CreationTime: "2018-07-23T05:59:26Z",
		UpdateTime:   "2018-07-23T05:59:26Z",
	})
	if err != nil {
		t.Log(err)
	}
}

func TestHarborClient_DeleteUser(t *testing.T) {
	client := NewHarborClient("localhost/", "admin", "Harbor12345")

	err := client.DeleteUser(6)
	if err != nil {
		t.Log(err)
	}
}

func TestHarborClient_UpdateUserPassword(t *testing.T) {
	client := NewHarborClient("localhost", "admin", "Harbor12345")

	err := client.UpdateUserPassword(&UpdateUser{
		UserId:      "7",
		OldPassWord: "Harbor12345",
		NewPassWord: "Harbor123456",
	})
	if err != nil {
		t.Log(err)
	}
}
