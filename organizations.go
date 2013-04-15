package gothub

import (
	"fmt"
)

type Organization struct {
	AvatarUrl string `json:"avatar_url"`
	EventsUrl string `json:"events_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	MembersUrl string `json:"members_url"`
	PublicMembersUrl string `json:"public_members_url"`
	ReposUrl string `json:"repos_url"`
	Url string `json:"url"`
}

// List all of the organizations the currently-authenticated user is a member of.
func (g *GitHub) Organizations() (orgs []Organization, err error) {
	err = g.callGithubApi("GET", "/user/orgs", &orgs)
	return
}

// List all of the organizations the user is a member of.
func (u *User) Organizations() (orgs []Organization, err error) {
	uri := fmt.Sprintf("/users/%s/orgs", u.Login)
	err = u.g.callGithubApi("GET", uri, &orgs)
	return
}