package ok

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

const (
	// Name is the name of the step.
	Name = "Get user"
	// Description is a brief description of the step.
	Description = "Retrieves the one user by the endpoint"
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
	// user is the user expects to be retrieved by the endpoint.
	user user.User
	// id is the id of the user to retrieve.
	id user.ID
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client and application base URL.
func NewStep(
	cli httpClient,
	baseURL string,
	id user.ID,
	user user.User,
) *Step {
	return &Step{
		cli:     cli,
		baseURL: baseURL,
		id:      id,
		user:    user,
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
