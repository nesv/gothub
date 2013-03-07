package gothub

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// Represents information about a user's account as defined here:
// http://developer.github.com/v3/users/#get-a-single-user
//
// Please take note that you cannot use the User struct to modify the details of a user's account.
// To do this, please look at the CurrentUser struct.
type User struct {
	Login       string    `json:"login"`
	Id          int       `json:"id"`
	AvatarUrl   string    `json:"avatar_url"`
	GravatarId  string    `json:"gravatar_id"`
	Url         string    `json:"url"`
	Name        string    `json:"name"`
	Company     string    `json:"company"`
	Blog        string    `json:"blog"`
	Location    string    `json:"location"`
	Email       string    `json:"email"`
	Hireable    bool      `json:"hireable"`
	Bio         string    `json:"bio"`
	PublicRepos int       `json:"public_repos"`
	PublicGists int       `json:"public_gists"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	HtmlUrl     string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
	Type        string    `json:"type"`
	g           *GitHub
}

func (u *User) Emails() (emails []string, err error) {
	response, err := call(u.g, "GET", "/user/emails")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	rs := strings.Trim(string(body), "[]")
	for _, email := range strings.Split(rs, ",") {
		emails = append(emails, strings.Trim(email, "\""))
	}
	return
}

// Returns the details of a single user, as specified by their "login".
//
// The term "login" is synonymous with "username":
//
//    https://github.com/<login>
func (g *GitHub) GetUser(login string) (*User, error) {
	var user User
	err := g.callGithubApi("GET", fmt.Sprintf("/users/%s", login), &user)
	if err != nil {
		return nil, err
	}
	user.g = g
	return &user, nil
}

// Returns the currently-authenticated user, as a pointer to a User struct.
func (g *GitHub) GetCurrentUser() (*User, error) {
	var user User
	err := g.callGithubApi("GET", "/user", &user)
	if err != nil {
		return nil, err
	}
	user.g = g
	return &user, nil
}
