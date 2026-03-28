package endpoint

import (
	"io"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

const (
	// Name is the name of the step.
	Name = "Create an event by endpoint"
	// Description is a brief description of the step.
	Description = "Creates the provived event by endpoint that to be found by fitler in the next steps"
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

// Step represents the HTTP create event step in the autograder process.
type Step struct {
	// cli is the HTTP client used to send requests.
	cli httpClient
	// baseURL is the base URL of the application.
	baseURL string
	// event is the event that has to be created by the target application.
	event event.Event
}

// NewStep creates a new Step instance
// with the provided HTTP client and application base URL.
func NewStep(
	cli httpClient,
	baseURL string,
	evnt event.Event,
) *Step {
	return &Step{
		cli:     cli,
		baseURL: baseURL,
		event:   evnt,
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
