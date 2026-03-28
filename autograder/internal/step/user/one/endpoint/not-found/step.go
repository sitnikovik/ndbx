package notfound

import (
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

const (
	// Name is the name of the step.
	Name = "Get non-existent user"
	// Description is a brief description of the step.
	Description = "Checks behavior of the application for trying to get non-existent user"
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
	// user is the user that is not exist in the database.
	user user.User
	// baseURL is the base URL of the application.
	baseURL string
}

// NewStep creates a new Step instance
// with the provided HTTP client, application base URL
// and the users that is not exist in the database.
func NewStep(
	cli httpClient,
	baseURL string,
	usr user.User,
) *Step {
	return &Step{
		cli:     cli,
		baseURL: baseURL,
		user:    usr,
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
