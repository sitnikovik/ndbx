package endpoint

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

const (
	// Name is the name of the step.
	Name = "Get list of events for the user"
	// Description is a brief description of the step.
	Description = "Checks for list of the events for the users and filters them by all parameters"
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
	// events is the list of events to be checked.
	events []event.Event
	// rq     body.Body
	rq body.Body
	// id is the id of the user's events to retrieve.
	id user.ID
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client, application base URL, user's id,
// request body and expected events.
func NewStep(
	cli httpClient,
	baseURL string,
	id user.ID,
	rq body.Body,
	events []event.Event,
) *Step {
	return &Step{
		cli:     cli,
		baseURL: baseURL,
		id:      id,
		rq:      rq,
		events:  events,
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
