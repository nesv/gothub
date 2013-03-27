package gothub

import (
	"log"
	"testing"
)

var (
	tgh *GitHub
)

func init() {
	username, password, err := getTestingCredentials()
	if err != nil {
		log.Fatal(err)
	}

	tgh, err = BasicLogin(username, password)
	if err != nil {
		log.Fatal("Failed to login to the Github")
	}
}

func TestSearchIssues(t *testing.T) {
	rs, err := tgh.SearchIssues(&IssuesSearchParam{
		Owner:      "nesv",
		Repository: "gothub",
		State:      "open",
		Keyword:    "endpoints",
	})
	if err != nil {
		t.Fatal(err)
	}

	for i, r := range rs {
		t.Logf("%d, %v", i, r)
	}
}

func TestSearchRepositories(t *testing.T) {
	rs, err := tgh.SearchRepositories(&RepositoriesSearchParam{
		Keyword: "whac a gopher",
	})
	if err != nil {
		t.Fatal(err)
	}

	for i, r := range rs {
		t.Logf("%d, %v", i, r)
	}
}

func TestSearchUsers(t *testing.T) {
	rs, err := tgh.SearchUsers(&UsersSearchParam{
		Keyword: "Homin Lee",
	})
	if err != nil {
		t.Fatal(err)
	}

	for i, r := range rs {
		t.Logf("%d, %v", i, r)
	}
}

func TestSearchEmail(t *testing.T) {
	r, err := tgh.SearchEmail(&EmailSearchParam{
		Email: "octocat@github.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", r)
}
