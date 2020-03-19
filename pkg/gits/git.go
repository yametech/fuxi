package gits

type Git interface {
	ListRepositories(org string) ([]*GitRepository, error)
	ListBranch(repoName string) ([]string, error)
}

// Repository represents a repository
type GitRepository struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Size        int    `json:"size"`
	HTMLURL     string `json:"html_url"`
	SSHURL      string `json:"ssh_url"`
	CloneURL    string `json:"clone_url"`
	OriginalURL string `json:"original_url"`
}

type GitArgs struct {
	Username string `json:"username"`
	ApiToken string `json:"apitoken"`
	Password string `json:"password,omitempty"`
	Url      string `json:"url"`
}
