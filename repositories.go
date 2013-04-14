package gothub

import (
	"fmt"
	"time"
)

type RepositoryOwner struct {
	Login      string `json:"login"`
	Id         int    `json:"id"`
	AvatarUrl  string `json:"avatar_url"`
	GravatarId string `json:"gravatar_id"`
	Url        string `json:"url"`
}

type Repository struct {
	Id               int             `json:"id"`
	Owner            RepositoryOwner `json:"owner"`
	Name             string          `json:"name"`
	FullName         string          `json:"full_name"`
	Description      string          `json:"description"`
	Private          bool            `json:"private"`
	Fork             bool            `json:"fork"`
	Url              string          `json:"url"`
	HtmlUrl          string          `json:"html_url"`
	CloneUrl         string          `json:"clone_url"`
	GitUrl           string          `json:"git_url"`
	SshUrl           string          `json:"ssh_url"`
	SvnUrl           string          `json:"svn_url"`
	MirrorUrl        string          `json:"mirror_url"`
	ArchiveUrl       string          `json:"archive_url"`
	AssigneesUrl     string          `json:"assignees_url"`
	BlobsUrl         string          `json:"blobs_url"`
	CollaboratorsUrl string          `json:"collaborators_url"`
	CommentsUrl      string          `json:"comments_url"`
	CompareUrl       string          `json:"compare_url"`
	ContentsUrl      string          `json:"contentsUrl"`
	ContributorsUrl  string          `json:"contributors_url"`
	Homepage         string          `json:"homepage"`
	Forks            int             `json:"forks"`
	ForksCount       int             `json:"forks_count"`
	Watchers         int             `json:"watchers"`
	WatchersCount    int             `json:"watchers_count"`
	Size             int             `json:"size"`
	MasterBranch     string          `json:"master_branch"`
	OpenIssues       string          `json:"open_issues"`
	PushedAt         time.Time       `json:"pushed_at"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

// Get the currently-authenticated user's repositories.
func (g GitHub) Repositories() (repositories []Repository, err error) {
	repositories = make([]Repository, DefaultPageSize)
	err = g.callGithubApi("GET", "/user/repos", &repositories)
	return
}

// Get the user's repositories.
func (u User) Repositories() (repositories []Repository, err error) {
	repositories = make([]Repository, DefaultPageSize)
	uri := fmt.Sprintf("/users/%s/repos", u.Login)
	err = u.g.callGithubApi("GET", uri, &repositories)
	return
}
