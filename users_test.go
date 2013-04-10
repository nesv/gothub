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

func TestGetFollowing(t *testing.T) {
	user, _ := tgh.GetCurrentUser()
	if following, err := user.GetFollowing(); err != nil {
		t.Error(err)
	} else {
		t.Logf("%s is following:", user.Login)
		for _, f := range following {
			t.Logf("%s", f.Login)
		}
	}
}

func TestIsFollowing(t *testing.T) {
	u := "octocat"
	if followingp, err := tgh.IsFollowing(u); err != nil {
		t.Error(err)
	} else if followingp {
		t.Logf("The current user is following %s", u)
	} else {
		t.Logf("The current user is not following %s", u)
	}
}

func TestFollow(t *testing.T) {
	u := "octocat"
	if err := tgh.Follow(u); err != nil {
		t.Error(err)
	} else {
		if following, err := tgh.IsFollowing(u); err != nil {
			t.Error(err)
		} else if following {
			t.Logf("You are now following %s", u)
		} else {
			t.Logf("Caches take time")
		}
	}
}

func TestUnfollow(t *testing.T) {
	u := "octocat"
	if err := tgh.Unfollow(u); err != nil {
		t.Error(err)
	} else {
		if following, err := tgh.IsFollowing(u); err != nil {
			t.Error(err)
		} else if !following {
			t.Logf("You are no longer following %s", u)
		} else {
			t.Logf("Caches take time")
		}
	}
}
