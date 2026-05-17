package endpoint

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/endpoint/expect"
)

// httpClient defines the interface for making HTTP requests.
type httpClient interface {
	// Get sends a GET request to the specified URL and returns the response.
	Get(url string) (*http.Response, error)
}

// Step represents the HTTP create event step in the autograder process.
type Step struct {
	// cli is the HTTP client used to send requests.
	cli httpClient
	// want holds the expectations we need to check in the Step.
	want expect.Expectations
	// user is the user's which events are expected to be retrieved.
	user user.User
	// desc is the description of the step.
	desc step.Desc
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client, application base URL, user,
// request body, expected events and additional functional options.
func NewStep(
	desc step.Desc,
	cli httpClient,
	baseURL string,
	usr user.User,
	want expect.Expectations,
) *Step {
	s := &Step{
		desc:    desc,
		cli:     cli,
		baseURL: baseURL,
		user:    usr,
		want:    want,
	}
	return s
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return s.desc.Title()
}

// Description returns a brief description of what the step does.
func (s *Step) Description() string {
	return s.desc.Description()
}
