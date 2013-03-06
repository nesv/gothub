package gothub

import (
	"fmt"
	"net/url"
	"time"
)

// Represents information about a user's account as defined here:
// http://developer.github.com/v3/users/#get-a-single-user
//
// Please take note that you cannot use the User struct to modify the details of a user's account.
// To do this, please look at the CurrentUser struct.
type User struct {
	Login       string `json:"login"`
	Id          int    `json:"id"`
	AvatarUrl   url.URL
	GravatarId  string `json:"gravatar_id"`
	Url         url.URL
	Name        string `json:"name"`
	Company     string `json:"company"`
	Blog        url.URL
	Location    string `json:"location"`
	Email       string `json:"email"`
	Hireable    bool   `json:"hireable"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	PublicGists int    `json:"public_gists"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	HtmlUrl     url.URL
	CreatedAt   time.Time
	Type        string `json:"type"`
}

// Returns the details of a single user, as specified by their "login".
//
// The term "login" is synonymous with "username":
//
//    https://github.com/<login>
func (g *GitHub) GetUser(login string) (*User, error) {
	var user User
	if err := g.callGithubApi("GET", fmt.Sprintf("/users/%s", login), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Get a list of every single user on GitHub.
func GetAllUsers() ([]*User, error) {
	return nil, nil
}

// Returns the currently-authenticated user, as a pointer to a User struct.
func (g *GitHub) GetCurrentUser() (*User, error) {
	var user User
	if err := g.callGithubApi("GET", "/user", &user); err != nil {
		return nil, err
	}
	return &user, nil
}
