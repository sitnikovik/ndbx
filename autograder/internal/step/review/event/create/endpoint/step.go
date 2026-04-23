package endpoint

import (
	"io"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/create/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
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
	// event is the event that has to be created by the target application.
	event event.Event
	// desc is the description of the step.
	desc step.Desc
	// rq is the request body to send.
	rq body.Body
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client and application base URL.
func NewStep(
	desc step.Desc,
	cli httpClient,
	baseURL string,
	ev event.Event,
	rq body.Body,
) *Step {
	return &Step{
		cli:     cli,
		desc:    desc,
		baseURL: baseURL,
		event:   ev,
		rq:      rq,
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
