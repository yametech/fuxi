package harbor

import "testing"

func TestHarborClient_ListProjectMembers(t *testing.T) {
	client := NewHarborClient("localhost", "admin", "Harbor12345")

	projects, err := client.ListProjectMembers(&MemberOptions{
		ProjectID:  6,
		EntityName: "",
	})
	if err != nil {
		t.Log(err)
	}
	t.Log(projects)
}

func TestHarborClient_CreateProjectMember(t *testing.T) {
	client := NewHarborClient("localhost", "admin", "Harbor12345")

	err := client.CreateProjectMember("6", &Member{
		RoleID: 1,
		MemberUser: MemberUser{
			UserID:   3,
			Username: "刘召",
		},
		MemberGroup: MemberGroup{
			ID:          0,
			GroupName:   "",
			GroupType:   0,
			LdapGroupDn: "",
		},
	})
	if err != nil {
		t.Log(err)
	}
}

func TestHarborClient_UpdateProjectMember(t *testing.T) {
	client := NewHarborClient("localhost", "admin", "Harbor12345")

	err := client.UpdateProjectMember(&MemberUpdateOption{
		RoleID: 1,
	}, "6", "10")
	if err != nil {
		t.Log(err)
	}
}

func TestHarborClient_DeleteProjectMember(t *testing.T) {
	client := NewHarborClient("localhost", "admin", "Harbor12345")

	err := client.DeleteProjectMember("6", "10")
	if err != nil {
		t.Log(err)
	}
}
