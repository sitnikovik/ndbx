package endpoint

import (
	"net/http"
)

const (
	// Name is the name of the preserve endpoint step.
	Name = "Preserve Endpoint Step"
	// Description is a brief explanation of what the preserve endpoint step does.
	Description = "Checks if the user's session is preserved by making an HTTP request" +
		" to the application's endpoint and verifying there is a session cookie in the response."
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
	httpcli httpClient,
	baseURL string,
) *Step {
	return &Step{
		cli:     httpcli,
		baseURL: baseURL,
	}
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return Name
}

// Description returns a brief explanation of what the step does.
func (s *Step) Description() string {
	return Description
}
