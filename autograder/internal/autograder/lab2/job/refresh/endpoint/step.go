package endpoint

import (
	"io"
	"net/http"
)

const (
	// Name is the name of the refresh session endpoint step.
	Name = "Refresh Session Endpoint Step"
	// Description is a brief explanation of what the refresh session endpoint step does.
	Description = "Checks if the application does not recreate a session " +
		"when the user already has a valid session cookie by making an HTTP request " +
		"to the application's endpoint and verifying there is a session cookie in the response."
)

// httpClient defines the interface for making HTTP requests.
type httpClient interface {
	// PostJSON sends a POST request with a JSON body
	// to the specified URL and returns the response.
	PostJSON(
		url string,
		body io.Reader,
	) (*http.Response, error)
}

// Step represents the HTTP check-in step in the autograder process.
type Step struct {
	// cli is the HTTP client used to send requests.
	cli httpClient
	// url is the endpoint to which the check-in POST request will be sent.
	url string
}

// NewStep creates a new Step instance
// with the provided HTTP client and application base URL.
func NewStep(
	cli httpClient,
	url string,
) *Step {
	return &Step{
		cli: cli,
		url: url,
	}
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return Name
}

// Description returns a brief description of what the step does.
func (s *Step) Description() string {
	return Description
}
