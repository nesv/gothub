package gothub

import (
	"errors"
	"os"
	"testing"
)

func getTestingCredentials() (username, password string, err error) {
	username = os.Getenv("GITHUB_USERNAME")
	password = os.Getenv("GITHUB_PASSWORD")

	if len(username) == 0 {
		err = errors.New("The GITHUB_USERNAME envvar is not set")
		return
	}

	if len(password) == 0 {
		err = errors.New("The GITHUB_PASSWORD envvar is not set")
		return
	}

	return
}

func TestGuest(t *testing.T) {
	if g, err := Guest(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("RateLimit-Limit: %d", g.RateLimit)
		t.Logf("RateLimit-Remaining: %d", g.RateLimitRemaining)
	}
}

func TestBasicAuth(t *testing.T) {
	username, password, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var g *GitHub
	if g, err = BasicLogin(username, password); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Authorization: %s", g.Authorization)
		t.Logf("RateLimit-Limit: %d", g.RateLimit)
		t.Logf("RateLimit-Remaining: %d", g.RateLimitRemaining)
	}
}
