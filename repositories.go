package gothub

import (
	"fmt"
	"time"
)

type RepositoryOwner struct {
	AvatarUrl        string `json:"avatar_url"`
	EventsUrl        string `json:"events_url"`
	FollowersUrl     string `json:"followers_url"`
	FollowingUrl     string `json:"following_url"`
	GistsUrl         string `json:"gists_url"`
	GravatarId       string `json:"gravatar_id"`
	HtmlUrl          string `json:"html_url"`
	Id               int    `json:"id"`
	Login            string `json:"login"`
	OrganizationsUrl string `json:"organizations_url"`
	ReposUrl         string `json:"repos_url"`
	StarredUrl       string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	Type             string `json:"type"`
	Url              string `json:"url"`
}

type RepositoryPermissions struct {
	Admin bool `json:"admin"`
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
}

type Repository struct {
	ArchiveUrl       string                `json:"archive_url"`
	AssigneesUrl     string                `json:"assignees_url"`
	BlobsUrl         string                `json:"blobs_url"`
	CloneUrl         string                `json:"clone_url"`
	CollaboratorsUrl string                `json:"collaborators_url"`
	CommentsUrl      string                `json:"comments_url"`
	CommitsUrl       string                `json:"comments_url"`
	CompareUrl       string                `json:"compare_url"`
	ContentsUrl      string                `json:"contentsUrl"`
	ContributorsUrl  string                `json:"contributors_url"`
	CreatedAt        time.Time             `json:"created_at"`
	DefaultBranch    string                `json:"default_branch"`
	Description      string                `json:"description"`
	DownloadsUrl     string                `json:"downloads_url"`
	EventsUrl        string                `json:"events_url"`
	Fork             bool                  `json:"fork"`
	Forks            int                   `json:"forks"`
	ForksCount       int                   `json:"forks_count"`
	ForksUrl         string                `json:"forks_url"`
	FullName         string                `json:"full_name"`
	GitCommitsUrl    string                `json:"git_commits_url"`
	GitRefsUrl       string                `json:"git_refs_url"`
	GitTagsUrl       string                `json:"git_tags_url"`
	GitUrl           string                `json:"git_url"`
	HasDownloads     bool                  `json:"has_downloads"`
	HasIssues        bool                  `json:"has_issues"`
	HasWiki          bool                  `json:"has_wiki"`
	Homepage         string                `json:"homepage"`
	HooksUrl         string                `json:"hooks_url"`
	HtmlUrl          string                `json:"html_url"`
	Id               int                   `json:"id"`
	IssueCommentUrl  string                `json:"issue_comment_url"`
	IssueEventsUrl   string                `json:"issue_events_url"`
	IssuesUrl        string                `json:"issues_url"`
	KeysUrl          string                `json:"keys_url"`
	LabelsUrl        string                `json:"labels_url"`
	Language         string                `json:"language"`
	LanguagesUrl     string                `json:"languages_url"`
	MasterBranch     string                `json:"master_branch"`
	MergesUrl        string                `json:"merges_url"`
	MilestonesUrl    string                `json:"milestones_url"`
	MirrorUrl        string                `json:"mirror_url"`
	Name             string                `json:"name"`
	NotificationsUrl string                `json:"notifications_url"`
	OpenIssues       int                   `json:"open_issues"`
	OpenIssuesCount  int                   `json:"open_issues_count"`
	Owner            RepositoryOwner       `json:"owner"`
	Permissions      RepositoryPermissions `json:"permissions"`
	Private          bool                  `json:"private"`
	PullsUrl         string                `json:"pulls_url"`
	PushedAt         time.Time             `json:"pushed_at"`
	Size             int                   `json:"size"`
	SshUrl           string                `json:"ssh_url"`
	StargazersUrl    string                `json:"stargazers_url"`
	StatusesUrl      string                `json:"statuses_url"`
	SubscribersUrl   string                `json:"subscribers_url"`
	SubscriptionUrl  string                `json:"subscription_url"`
	SvnUrl           string                `json:"svn_url"`
	TagsUrl          string                `json:"tags_url"`
	TeamsUrl         string                `json:"teams_url"`
	TreesUrl         string                `json:"trees_url"`
	UpdatedAt        time.Time             `json:"updated_at"`
	Url              string                `json:"url"`
	Watchers         int                   `json:"watchers"`
	WatchersCount    int                   `json:"watchers_count"`
}

// Get the currently-authenticated user's repositories.
func (g GitHub) Repositories() (repositories []Repository, err error) {
	repositories = make([]Repository, 0)
	err = g.callGithubApi("GET", "/user/repos", &repositories)
	return
}

// Get the user's repositories.
func (u User) Repositories() (repositories []Repository, err error) {
	repositories = make([]Repository, 0)
	uri := fmt.Sprintf("/users/%s/repos", u.Login)
	err = u.g.callGithubApi("GET", uri, &repositories)
	return
}
