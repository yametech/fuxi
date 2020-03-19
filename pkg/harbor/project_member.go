package harbor

import (
	"strconv"

	"github.com/pkg/errors"
)

var (
	CreateProjectMemberError = errors.New("create project member error")
)

//Member member
type Member struct {
	RoleID      int         `json:"role_id"`
	MemberUser  MemberUser  `json:"member_user"`
	MemberGroup MemberGroup `json:"member_group"`
}

//MemberUser a member user
type MemberUser struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

//MemberGroup a Member Group
type MemberGroup struct {
	ID          int    `json:"id"`
	GroupName   string `json:"group_name"`
	GroupType   int    `json:"group_type"`
	LdapGroupDn string `json:"ldap_group_dn"`
}

//MemberRes a Member Response
type MemberRes struct {
	ID         int    `json:"id"`
	ProjectID  int    `json:"project_id"`
	EntityName string `json:"entity_name"`
	RoleName   string `json:"role_name"`
	RoleID     int    `json:"role_id"`
	EntityID   int    `json:"entity_id"`
	EntityType string `json:"entity_type"`
}

//MemberOptions a Member Options
type MemberOptions struct {
	ProjectID  int    `json:"project_id"`
	EntityName string `json:"entity_name"`
}

//MemberUpdateOption a MemberUpdate Option
type MemberUpdateOption struct {
	RoleID int `json:"role_id"`
}

//ListProjects get all projectmembers
func (c *HarborClient) ListProjectMembers(opt *MemberOptions) ([]MemberRes, error) {
	var members []MemberRes
	id := strconv.Itoa(opt.ProjectID)
	_, err := c.Client.R().
		SetResult(&members).
		Get(GETPROJECT + id + "/members")
	if err != nil {
		return members, err
	}
	return members, nil
}

//CreateProjectMember create a  projectmember
func (c *HarborClient) CreateProjectMember(projectId string, member *Member) error {
	res, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(member).
		Post(GETPROJECT + projectId + "/members")
	if err != nil {
		return err
	}
	if res.StatusCode() == 201 {
		return err
	}
	return errors.WithStack(CreateProjectMemberError)
}

//CreateProjectMember update  a  projectmember
func (c *HarborClient) UpdateProjectMember(memberUpdateOption *MemberUpdateOption, projectId, mid string) error {
	res, err := c.Client.R().SetBody(memberUpdateOption).Put(GETPROJECT + projectId + "/members/" + mid)
	if err != nil {
		return err
	}
	if res.StatusCode() == 200 {
		return nil
	}
	return errors.Wrap(UpdateProjectError, string(res.Body()))
}

//DeleteProjectMember update  a  projectmember
func (c *HarborClient) DeleteProjectMember(projectId, mid string) error {
	res, err := c.Client.R().Delete(GETPROJECT + projectId + "/members/" + mid)
	if err != nil {
		return err
	}
	if res.StatusCode() == 200 {
		return nil
	}
	return errors.Wrap(UpdateProjectError, "update project member fail")
}
