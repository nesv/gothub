package gothub

import (
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
func GetUser(login string) (*User, error) {
	var user User
	return &user, nil
}

// Get a list of every single user on GitHub.
//
// Please use sparingly.
func GetAllUsers() ([]*User, error) {
	return nil, nil
}

// A special struct to represent the current user.
//
// The only difference between the CurrentUser and User structs, is that the CurrentUser struct
// has methods pinned to it, and supports modification.
type CurrentUser struct {
	User
	modified bool
}

// Returns the currently-authenticated user, as a pointer to a CurrentUser struct.
func (g *GitHub) GetCurrentUser() (*CurrentUser, error) {
	var user CurrentUser
	if err := g.callGithubApi("GET", "/user", user); err != nil {
		return nil, err
	}
	return &user, nil
}
