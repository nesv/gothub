package gothub

import (
	"log"
	"testing"
)

var (
	testKeyId   int
	testSshKeys []string
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

	testSshKeys = []string{
		"ssh-dss AAAAB3NzaC1kc3MAAACBANTHulCj21/003YdYOqn0mfXc2JtI26haO2z18HqdA4GM6GglTJNepRZnatxH+J7UeQGhA5nChOeBw/pGm3vQ3WxKbrwKl8V2Ag0IdEIRmpc5j3Dx6ihl0jc1D+veVz6xUrqOPzu7YPeDYweUZE6b4L2FQq0Q9QvoVRXlIw1w+9BAAAAFQD32crUBTOaLHglubAGrMpT2irF0QAAAIEAx9E3v1FWFKWXjf3fihBiMfXdON3aOGF1zsH78ZEwXsaxHS9TmuBBClYSSSDzkZPYr0B0lTJgSo6rh9wuIRZul+tKDiNvbND/zl9h1ib2tt3VbfDgJlBQ6NoFt1ZHYZggv7jPogVD+/vRmksjIHp0nejI+EqWB+33gRyge6qu7VsAAACAKO78TWWhCAsGdU2uoGsxlYt9Mj7wphjJxwPvY5RIpT2mfwf0UP0u4R8vospmu9xf3Kqvh4qCztIUIyVGANw55eCzTaKrKFOBkUJqQRKEcpeuePWDIy+MOFWgkFtDbPtVGaziVui5Ujy5anap8EBPb3bFt1cdJioxSLRSREnBMOo= GotHub test key",
		"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCcPo4eTziK61PW+mJVUmUa49r9V731tfwvw/u5LTJ4KZ03+lH5ypxcCQ/32FKvpKRPdLlkgOoj6WxgzwoscqORWrYxmQaOnKiCZuSzCO2ndPgv4/EHQz4VpcxJHJsKIUfeAqjfQDI2WG6LM3iRCc03aIHP/H92tNCX36gX2jyc16mrYjb1+8zMrDLOMv9mPSSjynXqCMoKP7IqsfHfRy+Pd+Knab3nb4VN1ERhSzBdb6Ly8RegZZG3HB4VMheHkld/PcCEky+wVbplA3pirbiL7eMehPCTz8t/cHhwJhoTFlsY3U0CV+7KO6sGPYdas3rgHw8QeyxbFGquwqleJrLr"}
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

func TestGetPublicKeys(t *testing.T) {
	user, _ := tgh.GetCurrentUser()
	keys, err := user.GetPublicKeys()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s has the following public keys:", user.Login)
		for _, k := range keys {
			t.Logf("%d\t%s", k.Id, k.Key)
		}
	}
}

func TestCurrentUserPublicKeys(t *testing.T) {
	if keys, err := tgh.PublicKeys(); err != nil {
		t.Error(err)
	} else {
		t.Logf("You have the following public keys:")
		for _, k := range keys {
			t.Logf("%d %s %s", k.Id, k.Title, k.Url)
		}
	}
}

func TestGetSinglePublicKey(t *testing.T) {
	keys, err := tgh.PublicKeys()
	if err != nil {
		t.Error("Could not fetch the current user's public keys")
		t.Error(err)
	} else {
		key, err := tgh.GetPublicKey(keys[0].Id)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Fetched public key %d, %s", key.Id, key.Title)
		}
	}
}

func TestAddPublicKey(t *testing.T) {
	title := "gothub test key"
	if newKeyId, err := tgh.AddPublicKey(title, testSshKeys[0]); err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("Created new key \"%s\": %d", title, newKeyId)
		testKeyId = newKeyId
	}
}

func TestUpdatePublicKey(t *testing.T) {
	title := "gothub test key"
	if testKeyId == 0 {
		t.Errorf("The test key ID was not set by the TestAddPublicKey test")
		return
	}

	updatedKey, err := tgh.UpdatePublicKey(testKeyId, testSshKeys[0], title)
	if err != nil {
		t.Errorf("Could not update key %d: %s", testKeyId, err)
	} else {
		t.Logf("Key %s successfully modified -> %s", updatedKey.Id, updatedKey.Title)
	}
}

func TestRemovePublicKey(t *testing.T) {
	if testKeyId == 0 {
		t.Errorf("The test key ID was not set by the TestAddPublicKey test")
		return
	}

	if err := tgh.RemovePublicKey(testKeyId); err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("Successfully removed public key %d", testKeyId)
	}
}
