package gothub

import (
	"log"
	"testing"
)

func init() {
	u, p, err := getTestingCredentials()
	if err != nil {
		log.Fatal(err)
	}

	tgh, err = BasicLogin(u, p)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetUser(t *testing.T) {
	if user, err := tgh.GetUser("nesv"); err != nil {
		t.Error(err)
	} else {
		t.Logf("ID: %d", user.Id)
		t.Logf("Login: %s", user.Login)
	}
}

func TestGetCurrentUser(t *testing.T) {
	user, err := tgh.GetCurrentUser()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("ID: %d", user.Id)
		t.Logf("Login: %s", user.Login)
	}
}

func TestUserEmails(t *testing.T) {
	emails, err := tgh.Emails()
	if err != nil {
		t.Error(err)
	}

	t.Logf("# emails: %d", len(emails))
	for i, email := range emails {
		t.Logf("Email #%d: %s", i+1, email)
	}
}

func TestGetFollowers(t *testing.T) {
	user, _ := tgh.GetCurrentUser()
	if followers, err := user.GetFollowers(); err != nil {
		t.Error(err)
	} else {
		t.Logf("The following users are following \"%s\":", user.Login)
		for _, follower := range followers {
			t.Logf("%s", follower.Login)
		}
	}
}
