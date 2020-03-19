package gits

import "code.gitea.io/sdk/gitea"

type GiteaClient struct {
	Client   *gitea.Client
	UserName string
}

//NewGiteaClient new a gitea client
func NewGiteaClient(gitArgs *GitArgs) *GiteaClient {
	client := gitea.NewClient(gitArgs.Url, gitArgs.ApiToken)
	return &GiteaClient{Client: client, UserName: gitArgs.Username}
}

//ListRepositories get current user  all repos by org
func (client *GiteaClient) ListRepositories(org string) ([]*GitRepository, error) {
	var gitRepos []*GitRepository
	if org == "" {
		repos, err := client.Client.ListMyRepos()
		if err != nil {
			return gitRepos, err
		}
		for _, repo := range repos {
			gitRepos = append(gitRepos, toGiteaRepo(repo.Name, repo))
		}
		return gitRepos, nil
	}
	repos, err := client.Client.ListOrgRepos(org)
	if err != nil {
		return gitRepos, err
	}
	for _, repo := range repos {
		gitRepos = append(gitRepos, toGiteaRepo(repo.Name, repo))
	}
	return gitRepos, nil
}

//ListBranch get a repo all Branch
func (client *GiteaClient) ListBranchs(repoName string) ([]string, error) {
	var branchArray []string
	branchs, err := client.Client.ListRepoBranches(client.UserName, repoName)
	if err != nil {
		return branchArray, err
	}
	for _, branch := range branchs {
		branchArray = append(branchArray, branch.Name)
	}
	return branchArray, nil
}

//toGiteaRepo Mapper GiteaRepo Entity
func toGiteaRepo(name string, repo *gitea.Repository) *GitRepository {
	return &GitRepository{
		Name:     name,
		CloneURL: repo.CloneURL,
		HTMLURL:  repo.HTMLURL,
		SSHURL:   repo.SSHURL,
	}
}
