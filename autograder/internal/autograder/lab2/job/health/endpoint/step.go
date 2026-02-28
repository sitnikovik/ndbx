package endpoint

import (
	"net/http"
)

const (
	// Name is the name of the health check job.
	Name = "Healthcheck Endpoint step"
	// Description is a brief description of the health check job.
	Description = "Checks the health of the application by making an HTTP request " +
		"to the application's health check endpoint and verifying there is no session cookie in the response."
)

// httpClient defines the interface for making HTTP requests.
type httpClient interface {
	// Get sends a GET request to the specified URL and returns the response.
	Get(url string) (*http.Response, error)
}

// Step represents the HTTP health check step in the autograder process.
type Step struct {
	// cli is the HTTP client used to send requests.
	cli httpClient
	// baseURL is the base URL of the application being checked.
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
