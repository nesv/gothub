package gothub

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

// A minimized form of the User struct, for holding information about a follower.
type Follower struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
}

// Gets a detailed list of a user's followers.
func (u User) GetFollowers() (followers []Follower, err error) {
	uri := fmt.Sprintf("/users/%s/followers", u.Login)
	fs := make([]Follower, 0)
	err = u.g.callGithubApi("GET", uri, &fs)
	if err != nil {
		return
	}
	/*
		if len(response.Header.Get("Link")) > 0 {
			// It would appear that we have to make several more requests, so as 
			// to get all of the user's followers (read: pagination).
			pages := parseLinkHeader(response)
			for _, page := range pages {
				url, _ := url.Parse(page.Url)
				f := make([]Follower, 0)
				response, err = u.g.callGithubApi("GET", url.RequestURI(), &f)
				fs = append(fs, f...)
			}
		}
	*/
	followers = fs
	return
}

// Gets a list of users the user is following.
func (u User) GetFollowing() (following []Follower, err error) {
	uri := fmt.Sprintf("/users/%s/following", u.Login)
	following = make([]Follower, 0)
	err = u.g.callGithubApi("GET", uri, &following)
	return
}

// Holds information about user's public SSH keys that they have provided to GitHub.
type PublicKey struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

// Gets the verified public SSH keys for a user.
func (u User) GetPublicKeys() (keys []PublicKey, err error) {
	uri := fmt.Sprintf("/users/%s/keys", u.Login)
	keys = make([]PublicKey, 0)
	err = u.g.callGithubApi("GET", uri, &keys)
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

// Returns a list of the email accounts associated with the currently-
// authenticated user.
func (g *GitHub) Emails() (emails []string, err error) {
	response, err := call(g, "GET", "/user/emails")
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

// Associate a list of emails with the currently-authenticated user's account.
func (g *GitHub) AddEmails(emails []string) (err error) {
	addresses := fmt.Sprintf("%v", emails)
	buf := bytes.NewBufferString(addresses)
	response, err := g.httpPost("/user/emails", nil, buf)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusCreated {
		e := "GitHub returned a %d status code; was expecting %d"
		err = errors.New(fmt.Sprintf(e, response.StatusCode, http.StatusCreated))
	}
	return
}

// Disassociate a list of emails from the currently-authenticated user's
// account.
func (g *GitHub) DeleteEmails(emails []string) (err error) {
	addresses := fmt.Sprintf("%v", emails)
	buf := bytes.NewBufferString(addresses)
	response, err := g.httpDelete("/user/emails", nil, buf)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusNoContent {
		e := "GitHub returned a %d status code; was expecting %d"
		err = errors.New(fmt.Sprintf(e, response.StatusCode, http.StatusNoContent))
	}
	return
}

// Check to see whether or not the current user `u` is following another user.
func (g GitHub) IsFollowing(anotherUser string) (following bool, err error) {
	uri := fmt.Sprintf("/user/following/%s", anotherUser)
	response, err := g.httpGet(uri, nil)
	if err != nil {
		return
	}
	switch response.StatusCode {
	case http.StatusNoContent:
		following = true
	case http.StatusNotFound:
		following = false
	}
	return
}

// Follow a user.
func (g GitHub) Follow(user string) (err error) {
	uri := fmt.Sprintf("/user/following/%s", user)
	response, err := g.httpPut(uri, nil, nil)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusNoContent {
		e := "Bad HTTP status; wanted %d got %d"
		err = errors.New(fmt.Sprintf(e, http.StatusNoContent, response.StatusCode))
	}

	return
}

// Unfollow a user.
func (g GitHub) Unfollow(user string) (err error) {
	uri := fmt.Sprintf("/user/following/%s", user)
	response, err := g.httpDelete(uri, nil, nil)
	if err != nil {
		return
	} else if response.StatusCode != http.StatusNoContent {
		e := "Bad HTTP status; wanted %d got %d"
		err = errors.New(fmt.Sprintf(e, http.StatusNoContent, response.StatusCode))
	}
	return
}

// Fetch a listing of the currently-authenticated user's public SSH keys.
func (g GitHub) PublicKeys() (keys []PublicKey, err error) {
	keys = make([]PublicKey, 0)
	err = g.callGithubApi("GET", "/user/keys", &keys)
	return
}

// Fetch a singular public SSH key.
func (g GitHub) GetPublicKey(id int) (key PublicKey, err error) {
	uri := fmt.Sprintf("/user/keys/%d", id)
	err = g.callGithubApi("GET", uri, &key)
	return
}
