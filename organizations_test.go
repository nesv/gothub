package gothub

import "testing"

func TestGithubOrganizations(t *testing.T) {
	orgs, err := tgh.Organizations()
	if err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("You are a member of the following %s organizations:", len(orgs))
		for _, org := range orgs {
			t.Logf("%d\t%s", org.Id, org.Login)
		}
	}
}

func TestUserOrganizations(t *testing.T) {
	user, err := tgh.GetUser("octocat")
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	orgs, err := user.Organizations()
	if err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("%s is a part of the following %d orgs:", user.Login, len(orgs))
		for _, org := range orgs {
			t.Logf("%d\t%s", org.Id, org.Login)
		}
	}
}

func TestGetOrganization(t *testing.T) {
	if org, err := tgh.GetOrganization("github"); err != nil {
		t.Errorf("%s", err)
	} else {
		t.Logf("Successfully fetched organization \"%s\"", org.Name)
		t.Logf("%#v", org)
	}
}
