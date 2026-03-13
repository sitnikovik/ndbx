package endpoint

import (
	"io"
	"net/http"
)

const (
	// Name is the name of the step.
	Name = "Auth Failed Endpoint Step"
	// Description is a brief description of the step.
	Description = "Checks that the authentication endpoint correctly handles failed authentication attempts" +
		" by sending a POST request with invalid credentials and validating the response"
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

// Step represents the HTTP authentication step in the autograder process.
type Step struct {
	// cli is the HTTP client used to send requests.
	cli httpClient
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client and application base URL.
func NewStep(
	cli httpClient,
	baseURL string,
) *Step {
	return &Step{
		cli:     cli,
		baseURL: baseURL,
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
