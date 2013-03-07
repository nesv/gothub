package gothub

import (
	"testing"
)

func TestUserEmails(t *testing.T) {
	u, p, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	g, err := BasicLogin(u, p)
	if err != nil {
		t.Fatal(err)
	}

	user, err := g.GetCurrentUser()
	if err != nil {
		t.Fatal(err)
	}

	emails, err := user.Emails()
	if err != nil {
		t.Error(err)
	}

	t.Logf("# emails: %d", len(emails))
	for i, email := range emails {
		t.Logf("Email #%d: %s", i, email)
	}
}

func TestGetUser(t *testing.T) {
	u, p, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}
	var g *GitHub
	g, err = BasicLogin(u, p)
	if err != nil {
		t.Fatal(err)
	}
	if user, err := g.GetUser(u); err != nil {
		t.Error(err)
	} else {
		t.Logf("ID: %d", user.Id)
		t.Logf("Login: %s", user.Login)
	}
}

func TestGetCurrentUser(t *testing.T) {
	username, password, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var g *GitHub
	g, err = BasicLogin(username, password)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Authorization: %s", g.Authorization)

	user, err := g.GetCurrentUser()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("ID: %d", user.Id)
		t.Logf("Login: %s", user.Login)
	}
}
