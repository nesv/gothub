package gothub

import "testing"

func TestRepositories(t *testing.T) {
	reposCount := 0
	repos, err := tgh.Repositories()
	reposCount += len(repos)
	if err != nil {
		t.Errorf("%s", err)
	} else {
		for _, repo := range repos {
			t.Logf("%s", repo.FullName)
			//t.Logf("%#v", repo)
		}
	}

	repos, err = tgh.Repositories(2)
	reposCount += len(repos)
	if err != nil {
		t.Errorf("%s", err)
	} else {
		for _, repo := range repos {
			t.Logf("%s", repo.FullName)
		}
		t.Logf("You have %d repositories:", reposCount)
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

	repos, err = user.Repositories(2)
	if err != nil {
		t.Errorf("%s", err)
	} else {
		for _, repo := range repos {
			t.Logf("%s", repo.FullName)
		}
	}
	return
}
