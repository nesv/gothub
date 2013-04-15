package gothub

import (
	"fmt"
	"time"
)

type OrganizationPlan struct {
	Name         string `json:"name"`
	PrivateRepos int    `json:"private_repos"`
	Space        int    `json:"space"`
}

type Organization struct {
	AvatarUrl         string           `json:"avatar_url"`
	BillingEmail      string           `json:"billing_email"`
	Blog              string           `json:"blog"`
	Collaborators     int              `json:"collaborators"`
	Company           string           `json:"company"`
	CreatedAt         time.Time        `json:"created_at"`
	DiskUsage         int              `json:"disk_usage"`
	Email             string           `json:"email"`
	EventsUrl         string           `json:"events_url"`
	Followers         int              `json:"followers"`
	Following         int              `json:"following"`
	HtmlUrl           string           `json:"html_url"`
	Id                int              `json:"id"`
	Location          string           `json:"location"`
	Login             string           `json:"login"`
	MembersUrl        string           `json:"members_url"`
	Name              string           `json:"name"`
	OwnedPrivateRepos int              `json:"owned_private_repos"`
	Plan              OrganizationPlan `json:"plan"`
	PrivateGists      int              `json:"private_gists"`
	PublicGists       int              `json:"public_gists"`
	PublicMembersUrl  string           `json:"public_members_url"`
	PublicRepos       int              `json:"public_repos"`
	ReposUrl          string           `json:"repos_url"`
	TotalPrivateRepos int              `json:"total_private_repos"`
	Type              string           `json:"type"`
	UpdatedAt         time.Time        `json:"updated_at"`
	Url               string           `json:"url"`
}

// Returns a complete Organization struct.
func (g *GitHub) GetOrganization(name string) (org *Organization, err error) {
	uri := fmt.Sprintf("/orgs/%s", name)
	err = g.callGithubApi("GET", uri, &org)
	return
}

// List all of the organizations the currently-authenticated user is a member of.
//
// Please be aware that this method does not return a complete Organization struct;
// for that, please refer to the (*GitHub).GetOrganization() method.
func (g *GitHub) Organizations() (orgs []Organization, err error) {
	err = g.callGithubApi("GET", "/user/orgs", &orgs)
	return
}

// List all of the organizations the user is a member of.
//
// Please be aware that this method does not return a complete Organization struct;
// for that, please refer to the (*GitHub).GetOrganization() method.
func (u *User) Organizations() (orgs []Organization, err error) {
	uri := fmt.Sprintf("/users/%s/orgs", u.Login)
	err = u.g.callGithubApi("GET", uri, &orgs)
	return
}
