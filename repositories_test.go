package gothub

import "testing"

func TestRepositories(t *testing.T) {
	repos, err := tgh.Repositories()
	if err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("You have %d repositories:", len(repos))
		for _, repo := range repos {
			t.Logf("%s", repo.FullName)
			//t.Logf("%#v", repo)
		}
	}
}

func TestUserRepositories(t *testing.T) {
	user, err := tgh.GetUser("octocat")
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	repos, err := user.Repositories()
	if err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("%s has the following repositories:", user.Login)
		for _, repo := range repos {
			t.Logf("%s", repo.FullName)
			//t.Logf("%#v", repo)
		}
	}
	return
}
