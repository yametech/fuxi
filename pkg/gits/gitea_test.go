package gits

import (
	"fmt"
	"testing"
)

func TestGiteaClient_ListRepositories(t *testing.T) {
	gitArgs := &GitArgs{
		Username: "withlin",
		ApiToken: "",
		Url:      "https://gitea.com",
	}
	giteaClient := NewGiteaClient(gitArgs)
	a, err := giteaClient.ListRepositories("")
	fmt.Println(err)
	fmt.Println(a)

}

func TestGiteaClient_ListBranch(t *testing.T) {
	gitArgs := &GitArgs{
		Username: "withlin",
		ApiToken: "",
		Url:      "https://gitea.com",
	}
	giteaClient := NewGiteaClient(gitArgs)
	a, err := giteaClient.ListBranches("skp")
	fmt.Println(err)
	fmt.Println(a)
}
