package endpoint

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/list/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

const (
	// Name is the name of the step.
	Name = "Get list of users by filter"
	// Description is a brief description of the step.
	Description = "Gets the filtered list of the users stored in the application"
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
	// users is the users expects to be retrieved by the endpoint.
	users []user.User
	// rq is the request body of the endpoint to filter the users.
	rq body.Body
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client and application base URL.
func NewStep(
	cli httpClient,
	baseURL string,
	rq body.Body,
	users []user.User,
) *Step {
	return &Step{
		cli:     cli,
		baseURL: baseURL,
		users:   users,
		rq:      rq,
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
