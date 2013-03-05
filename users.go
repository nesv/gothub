package gothub

import (
	"net/url"
	"time"
	"errors"
)

// Represents information about a user's account as defined here:
// http://developer.github.com/v3/users/#get-a-single-user
//
// Please take note that you cannot use the User struct to modify the details of a user's account.
// To do this, please look at the CurrentUser struct.
type User struct {
	Login string
	Id int
	AvatarUrl url.URL
	GravatarId string
	Url url.URL
	Name string
	Company string
	Blog url.URL
	Location string
	Email string
	Hireable bool
	Bio string
	PublicRepos int
	PublicGists int
	Followers int
	Following int
	HtmlUrl url.URL
	CreatedAt time.Time
	Type string
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
	var users 
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
func GetCurrentUser() (*CurrentUser, error) {
	var user User
	return &user, nil
}