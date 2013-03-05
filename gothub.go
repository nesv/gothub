package gothub

import (
	"fmt"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"errors"
	"io/ioutil"
)

// The HTTP host that we will hit to use the GitHub API.
const (
	GitHubUrl string = "https://api.github.com"
	ErrRateLimitReached = errors.New("Rate limit reached")
)

// The GitHub struct represents an active session to the GitHub API.
type GitHub struct {
	httpClient *http.Client
	Authorization string
	rateLimit int
	rateLimitRemaining int
}

func hashAuth(u, p string) string {
	var a = fmt.Sprintf("%s:%s", u, p)
	return base64.StdEncoding.EncodeToString([]byte(a))
}

// Log in to GitHub using basic, username/password authentication.
func BasicLogin(username, password string) (*GitHub, error) {
	// Format and Base64-encode the provided username and password, in preparation for basic
	// HTTP auth.
	authorization := fmt.Sprintf("Basic %s", hashAuth(username, password))
	request, err := http.NewRequest("GET", GitHubUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set the Authorization header.
	request.Header.Set("Authorization", authorization)

	// Create a new HTTP client (which we will eventually provide to the GitHub struct), for
	// issuing the above HTTP request, and future requests.
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		// Yaaaaaaay!
		ratelimit := response.Header.Get("X-RateLimit-Limit")
		remaining := response.Header.Get("X-RateLimit-Remaining")
		return &GitHub{httpClient: client, Authorization: authorization, rateLimit: int(ratelimit), rateLimitRemaining: int(remaining)}, nil
	}

	// Should we get here, the basic authentication request failed.
	e := "Authorization failed with HTTP code: %d"
	return nil, errors.New(fmt.Sprintf(e, response.StatusCode))
}

func (g *GitHub) callGithubApi(method, uri string, rs interface{}) (err error) {
	if g.rateLimitRemaining == 0 {
		err = ErrRateLimitReached
		return
	}

	url := fmt.Sprintf("%s%s", GitHubUrl, uri)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	request.Header.Set("Authorization", g.Authorization)

	// Fire off the request.
	response, err := g.httpClient.Do(request)
	if err != nil {
		return
	}

	// Update the rate limits
	limit := response.Header.Get("X-RateLimit-Limit")
	g.rateLimit = int(limit)

	remaining := response.Header.Get("X-RateLimit-Remaining")
	g.rateLimitRemaining = int(remaining)
	

	// Now, marshal the HTTP response (if it was successful) onto the
	// struct provided by `rs`
	if response.StatusCode != http.StatusOK {
		e := "GitHub API responded with HTTP %d"
		err = errors.New(fmt.Sprintf(e, response.StatusCode))
		return
	}

	// Check to make sure we actually got JSON back.
	switch response.Header.Get("Content-Type") {
	case "application/json":
		fallthrough
	case "application/json; charset=utf-8":
		var js []byte
		js, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		} else {
			err = json.Unmarshal(js, &rs)
		}
	}

	return
}