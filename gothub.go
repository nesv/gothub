package gothub

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// The HTTP host that we will hit to use the GitHub API.
const (
	GitHubUrl string = "https://api.github.com"
)

var (
	ErrRateLimitReached = errors.New("Rate limit reached")
	ErrNoJSON           = errors.New("GitHub did not return a JSON response")
)

// The GitHub struct represents an active session to the GitHub API.
type GitHub struct {
	httpClient         *http.Client
	Authorization      string
	RateLimit          int
	RateLimitRemaining int
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
		ratelimit, err := strconv.Atoi(response.Header.Get("X-RateLimit-Limit"))
		if err != nil {
			return nil, err
		}

		remaining, err := strconv.Atoi(response.Header.Get("X-RateLimit-Remaining"))
		if err != nil {
			return nil, err
		}

		return &GitHub{httpClient: client, Authorization: authorization,
			RateLimit: ratelimit, RateLimitRemaining: remaining}, nil
	}

	// Should we get here, the basic authentication request failed.
	e := "Authorization failed with HTTP code: %d"
	return nil, errors.New(fmt.Sprintf(e, response.StatusCode))
}

// Updates the call limit rates in the GitHub struct.
func (g *GitHub) updateRates(r *http.Response) (err error) {
	limit, err := strconv.Atoi(r.Header.Get("X-RateLimit-Limit"))
	if err != nil {
		return
	}
	g.RateLimit = limit

	remaining, err := strconv.Atoi(r.Header.Get("X-RateLimit-Remaining"))
	if err != nil {
		return
	}
	g.RateLimitRemaining = remaining
	return
}

// Calls the GitHub API and returns the raw, HTTP response body.
func call(g *GitHub, method, uri string) (response *http.Response, err error) {
	if g.RateLimitRemaining == 0 {
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
	response, err = g.httpClient.Do(request)
	if err != nil {
		return
	}

	// Update the call rates
	g.updateRates(response)

	// Check to make sure the API came back with an appropriate HTTP status
	// code, depending on the request method
	switch method {
	case "GET":
		if response.StatusCode != http.StatusOK {
			e := "GitHub API responded with HTTP %d"
			err = errors.New(fmt.Sprintf(e, response.StatusCode))
		}

	case "POST":
		switch response.StatusCode {
		case http.StatusCreated:
			return
		}

	case "DELETE":
		switch response.StatusCode {
		case http.StatusNoContent:
			return
		}
	}

	return
}

// Calls the GitHub API, but will unmarshal a JSON response to the struct
// provided to `rs`.
func (g *GitHub) callGithubApi(method, uri string, rs interface{}) error {
	response, err := call(g, method, uri)
	if err != nil {
		return err
	}

	// Check to make sure we actually got JSON back.
	switch response.Header.Get("Content-Type") {
	case "application/json":
		fallthrough
	case "application/json; charset=utf-8":
		var js []byte
		js, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(js, rs)
		}
	default:
		err = ErrNoJSON
	}

	return err
}
