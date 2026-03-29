package endpoint

import (
	"io"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

const (
	// Name is the name of the step.
	Name = "Update Event Not Found"
	// Description is a brief description of the step.
	Description = "Checks the app behavior for someone user tries to change the event that does not relate for him"
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
	// event is the event to update.
	event event.Event
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client, application base URL and event to update.
func NewStep(
	cli httpClient,
	baseURL string,
	ev event.Event,
) *Step {
	return &Step{
		cli:     cli,
		baseURL: baseURL,
		event:   ev,
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
