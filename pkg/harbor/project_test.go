package harbor

import (
	"testing"
)

func TestHarborClient_ListProjects(t *testing.T) {
	client := NewHarborClient("localhost/", "admin", "Harbor12345")

	projects, err := client.ListProjects(&ProjectOption{
		Name:     "",
		Public:   "",
		Owner:    "",
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		t.Log(err)
	}
	t.Log(projects)
}

func TestHarborClient_CreateProject(t *testing.T) {
	client := NewHarborClient("localhost", "admin", "Harbor12345")

	err := client.CreateProject(&Project{
		CountLimit:   110,
		ProjectName:  "aaa",
		CveWhitelist: CveWhitelist{},
		StorageLimit: 10,
		Metadata: Metadata{
			Public:               "true",
			EnableContentTrust:   "false",
			PreventVul:           "false",
			Severity:             "High",
			AutoScan:             "false",
			ReuseSysCveWhitelist: "false",
		},
	})
	if err != nil {
		t.Log(err)
	}
}

func TestHarborClient_UpdateProject(t *testing.T) {
	client := NewHarborClient("localhost", "admin", "Harbor12345")

	err := client.UpdateProject(&UpdateProject{
		ProjectId:    0,
		CountLimit:   0,
		ProjectName:  "",
		CveWhitelist: CveWhitelist{},
		StorageLimit: 0,
		Metadata:     Metadata{},
	})
	if err != nil {
		t.Log(err)
	}
}

func TestHarborClient_DeleteProject(t *testing.T) {
	client := NewHarborClient("localhost", "admin", "Harbor12345")

	err := client.DeleteProject(7)
	if err != nil {
		t.Log(err)
	}
}
