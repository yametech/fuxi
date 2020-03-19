package harbor

import (
	"strconv"

	"github.com/pkg/errors"
)

const (
	GETPROJECT = "api/projects/"
)

var (
	CreateProjectError = errors.New("create a project  error")
	UpdateProjectError = errors.New("update a project  error")
	DeleteProjectError = errors.New("delete a project  error")
)

//ProjectListRes  a ProjectListRes
type ProjectListRes struct {
	ProjectID         int          `json:"project_id"`
	OwnerID           int          `json:"owner_id"`
	Name              string       `json:"name"`
	CreationTime      string       `json:"creation_time"`
	UpdateTime        string       `json:"update_time"`
	Deleted           bool         `json:"deleted"`
	OwnerName         string       `json:"owner_name"`
	Togglable         bool         `json:"togglable"`
	CurrentUserRoleID int          `json:"current_user_role_id"`
	RepoCount         int          `json:"repo_count"`
	ChartCount        int          `json:"chart_count"`
	Metadata          Metadata     `json:"metadata"`
	CveWhitelist      CveWhitelist `json:"cve_whitelist"`
}

//CveWhitelist CveWhitelist
type CveWhitelist struct {
	ID        int `json:"id"`
	ProjectID int `json:"project_id"`
	ExpiresAt int `json:"expires_at"`
	Items     []struct {
		CveID string `json:"cve_id"`
	} `json:"items"`
}

//Metadata Metadata
type Metadata struct {
	Public               string `json:"public"`
	EnableContentTrust   string `json:"enable_content_trust"`
	PreventVul           string `json:"prevent_vul"`
	Severity             string `json:"severity"`
	AutoScan             string `json:"auto_scan"`
	ReuseSysCveWhitelist string `json:"reuse_sys_cve_whitelist"`
}

//Project  a Project entity
type Project struct {
	CountLimit   int          `json:"count_limit"`
	ProjectName  string       `json:"project_name"`
	CveWhitelist CveWhitelist `json:"cve_whitelist"`
	StorageLimit int          `json:"storage_limit"`
	Metadata     Metadata     `json:"metadata"`
}

//UpdateProject a UpdateProject entity
type UpdateProject struct {
	ProjectId    int          `json:"project_id"`
	CountLimit   int          `json:"count_limit"`
	ProjectName  string       `json:"project_name"`
	CveWhitelist CveWhitelist `json:"cve_whitelist"`
	StorageLimit int          `json:"storage_limit"`
	Metadata     Metadata     `json:"metadata"`
}

//ProjectOption  ProjectOption
type ProjectOption struct {
	Name string `json:"name"`
	// NOTE:
	// 这里将 public 的类型从 bool 变更为 string ，因为bool 类型只有 true 和 false 二值语义，而实际使用中需要第三种语义
	// 1. 若为 true 则仅返回 public 项目；
	// 2. 若为 false 则仅返回 private 项目；
	// 3. 若不指定 public 参数，则应该同时返回 public 和 private 项目；
	Public string `json:"public"`
	// FIXME:
	// harbor 中基于 owner 过滤的功能似乎存在问题；
	Owner    string `json:"owner"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

//todo: 考虑做成可选参数，变成默认值，就取所有的仓储
//ListProjects get all project
func (c *HarborClient) ListProjects(opt *ProjectOption) ([]*ProjectListRes, error) {
	var projects []*ProjectListRes
	_, err := c.Client.R().SetPathParams(
		map[string]string{
			"name":      opt.Name,
			"public":    opt.Public,
			"owner":     opt.Owner,
			"page":      strconv.Itoa(opt.Page),
			"page_size": strconv.Itoa(opt.PageSize),
		}).
		SetResult(&projects).Get(GETPROJECT)
	if err != nil {
		return projects, err
	}
	return projects, nil
}

//CreateProject create a project
func (c *HarborClient) CreateProject(project *Project) error {
	res, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(project).
		Post(GETPROJECT)
	if err != nil {
		return nil
	}
	if res.StatusCode() == 201 {
		return nil
	}
	return errors.WithStack(CreateProjectError)
}

//UpdateProject  update a project
func (c *HarborClient) UpdateProject(opt *UpdateProject) error {
	id := strconv.Itoa(opt.ProjectId)
	res, err := c.Client.R().
		SetBody(opt).Put(GETPROJECT + id)
	if err != nil {
		return err
	}
	if res.StatusCode() == 200 {
		return nil
	}
	return errors.Wrap(UpdateProjectError, string(res.Body()))
}

//UpdateProject  delete a project
func (c *HarborClient) DeleteProject(projectId int) error {
	id := strconv.Itoa(projectId)
	res, err := c.Client.R().Delete(GETPROJECT + id)
	if err != nil {
		return err
	}
	if res.StatusCode() == 200 {
		return nil
	}
	return errors.Wrap(DeleteProjectError, string(res.Body()))
}
