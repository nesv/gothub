package gothub

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

var (
	ErrSearchInvalidParam = errors.New("Invalid parameter for search")
	ErrSearchOmitParam    = errors.New("Omit mandatory paramter for search")
)

// Types which represents params and result for search issues
// http://developer.github.com/v3/search/#search-issues
type IssuesSearchParam struct {
	Owner      string
	Repository string
	State      string // open or close
	Keyword    string

	result issuesSearchResult
}

func (p *IssuesSearchParam) addr() (string, error) {
	if len(p.Owner) == 0 || len(p.Repository) == 0 ||
		len(p.State) == 0 || len(p.Keyword) == 0 {
		return "", ErrSearchOmitParam
	}

	r := "/legacy/issues/search/"
	r += p.Owner
	r += ("/" + p.Repository)
	r += ("/" + p.State)
	r += ("/" + url.QueryEscape(p.Keyword))
	return r, nil
}

type issuesSearchResult struct {
	Issues []*IssuesSearchResult
}

type IssuesSearchResult struct {
	GravatarId string    `json:"gravatar_id"`
	Position   int       `json:"position"`
	Number     int       `json:"number"`
	Votes      int       `json:"votes"`
	CreatedAt  time.Time `json:"created_at"`
	Comments   int       `json:"comments"`
	Body       string    `json:"body"`
	Title      string    `json:"title"`
	UpdatedAt  time.Time `json:"updated_at"`
	HtmlUrl    string    `json:"html_url"`
	User       string    `json:"user"`
	Labels     []string  `json:"labels"`
	State      string    `json:"state"`
}

// Types which represents params and result for search repositories
// http://developer.github.com/v3/search/#search-repositories
type RepositoriesSearchParam struct {
	Keyword   string
	Language  string // optional:
	StartPage uint   // optional:
	Sort      string // optional: stars, forks or update
	Order     string // optional: asc or desc

	result repositoriesSearchResult
}

func (p *RepositoriesSearchParam) addr() (string, error) {
	if len(p.Keyword) == 0 {
		return "", ErrSearchOmitParam
	}

	r := "/legacy/repos/search/"
	r += url.QueryEscape(p.Keyword)
	if len(p.Language) != 0 {
		r += ("/" + p.Language)
	}
	if p.StartPage != 0 {
		r += fmt.Sprintf("/%d", p.StartPage)
	}
	if len(p.Sort) != 0 {
		r += ("/" + p.Sort)
	}
	if len(p.Order) != 0 {
		r += ("/" + p.Order)
	}
	return r, nil
}

type repositoriesSearchResult struct {
	Repositories []*RepositoriesSearchResult
}

type RepositoriesSearchResult struct {
	Type         string    `json:"type"`
	Created      time.Time `json:"created"`
	Watchers     int       `json:"watchers"`
	HasDownloads bool      `json:"has_downloads"`
	Username     string    `json:"username"`
	Homepage     string    `json:"homepage"`
	Url          string    `json:"url"`
	Fork         bool      `json:"fork"`
	HasIssues    bool      `json:"has_issues"`
	HasWiki      bool      `json:"has_wiki"`
	Forks        int       `json:"forks"`
	Size         int       `json:"size"`
	Private      bool      `json:"private"`
	Followers    int       `json:"followers"`
	Name         string    `json:"name"`
	Owner        string    `json:"owner"`
	OpenIssues   int       `json:"open_issues"`
	PushedAt     time.Time `json:"pushed_at"`
	Score        float32   `json:"score"`
	Pushed       time.Time `json:"pushed"`
	Description  string    `json:"description"`
	Language     string    `json:"language"`
	CreatedAt    time.Time `json:"created_at"`
}

// Types which represents params and result for search users
// http://developer.github.com/v3/search/#search-users
type UsersSearchParam struct {
	Keyword   string
	StartPage uint   // optional:
	Sort      string // optional: start, forks or update
	Order     string // optional: asc or desc

	result usersSearchResult
}

func (p *UsersSearchParam) addr() (string, error) {
	if len(p.Keyword) == 0 {
		return "", ErrSearchOmitParam
	}

	r := "/legacy/user/search/"
	r += url.QueryEscape(p.Keyword)
	if p.StartPage != 0 {
		r += fmt.Sprintf("/%d", p.StartPage)
	}
	if len(p.Sort) != 0 {
		r += ("/" + p.Sort)
	}
	if len(p.Order) != 0 {
		r += ("/" + p.Order)
	}
	return r, nil
}

type usersSearchResult struct {
	Users []*UsersSearchResult
}

type UsersSearchResult struct {
	GravatarId      string    `json:"gravatar_id"`
	Name            string    `json:"name"`
	CreatedAt       time.Time `json:"created_at"`
	Location        string    `json:"location"`
	PublicRepoCount int       `json:"public_repo_count"`
	Followers       int       `json:"followers"`
	Language        string    `json:"language"`
	Fullname        string    `json:"fullname"`
	Username        string    `json:"username"`
	Id              string    `json:"id"`
	Repos           int       `json:"repos"`
	Type            string    `json:"type"`
	FollowersCount  int       `json:"followers_count"`
	Login           string    `json:"login"`
	Score           float32   `json:"score"`
	Created         time.Time `json:"created"`
}

// Types which represents result for search email
// http://developer.github.com/v3/search/#email-users
type EmailSearchParam struct {
	Email string

	result emailSearchResult
}

func (p *EmailSearchParam) addr() (string, error) {
	if len(p.Email) == 0 {
		return "", ErrSearchOmitParam
	}
	if !strings.Contains(p.Email, "@") {
		return "", ErrSearchInvalidParam
	}

	return "/legacy/user/email/" + p.Email, nil
}

type emailSearchResult struct {
	User *EmailSearchResult
}

type EmailSearchResult struct {
	PublicRepoCount int       `json:"public_repo_count"`
	PublicGistCount int       `json:"public_gist_count"`
	FollowersCount  int       `json:"followers_count"`
	FollowingCount  int       `json:"following_count"`
	Created         time.Time `json:"created"`
	CreatedAt       time.Time `json:"created_at"`
	Name            string    `json:"name"`
	Company         string    `json:"company"`
	BLog            string    `json:"blog"`
	Location        string    `json:"location"`
	Email           string    `json:"email"`
	Id              int       `json:"id"`
	Login           string    `json:"login"`
	Type            string    `json:"type"`
	GravatarId      string    `json:"gravatar_id"`
}

// Find issues by state and keyword.
func (g *GitHub) SearchIssues(p *IssuesSearchParam) ([]*IssuesSearchResult, error) {
	addr, err := p.addr()
	if err != nil {
		return nil, err
	}
	if err := g.callGithubApi("GET", addr, &p.result); err != nil {
		return nil, err
	}
	return p.result.Issues, nil
}

// Find repositories by keyword.
// This method returns up to 100 results per page and
// pages can be fetched using the start_page parameter.
func (g *GitHub) SearchRepositories(p *RepositoriesSearchParam) ([]*RepositoriesSearchResult, error) {
	addr, err := p.addr()
	if err != nil {
		return nil, err
	}
	if err := g.callGithubApi("GET", addr, &p.result); err != nil {
		return nil, err
	}
	return p.result.Repositories, nil
}

// Find users by state and keyword.
func (g *GitHub) SearchUsers(p *UsersSearchParam) ([]*UsersSearchResult, error) {
	addr, err := p.addr()
	if err != nil {
		return nil, err
	}
	if err := g.callGithubApi("GET", addr, &p.result); err != nil {
		return nil, err
	}
	return p.result.Users, nil
}

// Find user by email address.
// This API call is added for compatibility reasons only.
// There’s no guarantee that full email searches will always be available.
func (g *GitHub) SearchEmail(p *EmailSearchParam) (*EmailSearchResult, error) {
	addr, err := p.addr()
	if err != nil {
		return nil, err
	}
	if err := g.callGithubApi("GET", addr, &p.result); err != nil {
		return nil, err
	}
	return p.result.User, nil
}
