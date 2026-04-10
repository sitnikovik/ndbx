package endpoint

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/events/list/expect"
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
	// expect is the expectations to check.
	expect expect.Expectations
	// rq is the request body to send.
	rq body.Body
	// desc is the description of the step.
	desc step.Desc
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client and application base URL.
func NewStep(
	desc step.Desc,
	cli httpClient,
	baseURL string,
	rq body.Body,
	want expect.Expectations,
) *Step {
	return &Step{
		cli:     cli,
		expect:  want,
		rq:      rq,
		desc:    desc,
		baseURL: baseURL,
	}
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return s.desc.Title()
}

// Description returns a brief description of what the step does.
func (s *Step) Description() string {
	return s.desc.Description()
}
