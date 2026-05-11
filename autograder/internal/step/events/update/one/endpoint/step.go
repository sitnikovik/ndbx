package endpoint

import (
	"io"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// httpClient defines the interface for making HTTP requests.
type httpClient interface {
	// Patch sends a PATCH request with a JSON body
	// to the specified URL and returns the response.
	Patch(
		url string,
		body io.Reader,
	) (*http.Response, error)
}

// Step represents the HTTP create event step in the autograder process.
type Step struct {
	// cli is the HTTP client used to send requests.
	cli httpClient
	// baseURL is the base URL of the application.
	baseURL string
	// event is the event that has to be created by the target application.
	event event.Event
	// want is the expectations to check.
	want expectation.Expectations
	// desc is the description of the step.
	desc step.Desc
}

// NewStep creates a new Step instance
// with the provided HTTP client and application base URL,
// event to create and expectations to check.
func NewStep(
	desc step.Desc,
	cli httpClient,
	baseURL string,
	evnt event.Event,
	want expectation.Expectations,
) *Step {
	return &Step{
		desc:    desc,
		cli:     cli,
		baseURL: baseURL,
		event:   evnt,
		want:    want,
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
