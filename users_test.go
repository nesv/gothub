package gothub

import (
	"testing"
)

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
	
	var user *CurrentUser
	user, err = g.GetCurrentUser()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("ID: %d", user.Id)
		t.Logf("Login: %s", user.Login)
	}
}