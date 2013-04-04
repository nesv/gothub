package gothub

import (
	"net/http"
	"net/url"
	"testing"
)

var (
	testUrl string = `https://api.github.com/users/bmizerany/followers`
)

func TestParseLinkHeader(t *testing.T) {
	response, err := http.Get(testUrl)
	if err != nil {
		t.Fatal(err)
		return
	}

	pages := parseLinkHeader(response)
	for _, page := range pages {
		// Make sure the Url field parses successfully.
		_, err := url.Parse(page.Url)
		if err != nil {
			t.Error(err)
		}

		// Make sure the Rel field is acceptable.
		switch page.Rel {
		case "prev", "first", "last", "next":
			t.Log("rel ok")
		default:
			t.Error("rel fail")
		}
		t.Logf("%+v", page)
	}
	return
}
