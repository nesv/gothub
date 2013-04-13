package gothub

import (
	"time"
	"fmt"
)

type Repository struct {
	Id int `json:"id"`
	Owner User `json:"owner"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Description string `json:"description"`
	Private bool `json:"private"`
	Fork bool `json:"fork"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	CloneUrl string `json:"clone_url"`
	GitUrl string `json:"git_url"`
	SshUrl string `json:"ssh_url"`
	SvnUrl string `json:"svn_url"`
	MirrorUrl string `json:"mirror_url"`
	Homepage string `json:"homepage"`
	Forks int `json:"forks"`
	ForksCount int `json:"forks_count"`
	Watchers int `json:"watchers"`
	WatchersCount int `json:"watchers_count"`
	Size int `json:"size"`
	MasterBranch string `json:"master_branch"`
	OpenIssues string `json:"open_issues"`
	PushedAt time.Time `json:"pushed_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Get the currently-authenticated user's repositories.
func (g GitHub) Repositories() (repositories []Repository, err error) {
	err = g.callGithubApi("GET", "/user/repos", &repositories)
	return
}

// Get the user's repositories.
func (u User) Repositories() (repositories []Repository, err error) {
	uri := fmt.Sprintf("/users/%s/repos", u.Login)
	err = u.g.callGithubApi("GET", uri, &repositories)
	return
}