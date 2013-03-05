package gothub

import (
	"fmt"
	"encoding/base64"
	"net/http"
	"errors"
)

// The HTTP host that we will hit to use the GitHub API.
const GitHubUrl string = "https://api.github.com"

// The GitHub struct represents an active session to the GitHub API.
type GitHub struct {
	httpClient &http.Client
	authorization string
}

// Log in to GitHub using basic, username/password authentication.
func BasicLogin(username, password string) (*GitHub, error) {
	// Format and Base64-encode the provided username and password, in preparation for basic
	// HTTP auth.
	basic := fmt.Sprintf("%s:%s", username, password)
	authorization := base64.StdEncoding.EncodeToString([]byte(basic))
	request, err := http.NewRequest("GET", GitHubUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set the Authorization header.
	request.Header.Set("Authorization", authorization)

	// Create a new HTTP client (which we will eventually provide to the GitHub struct), for
	// issuing the above HTTP request, and future requests.
	client := &http.Client{}
	response, err := client.Do(&request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		// Yaaaaaaay!
		return &GitHub{httpClient: client, authorization: authorization}, nil
	}

	// Should we get here, the basic authentication request failed.
	e := "Authorization failed with HTTP code: %d"
	return nil, errors.New(fmt.Sprintf(e, response.StatusCode))
}