package endpoint

import (
	"io"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

const (
	// Name is the name of the step.
	Name = "Auth login"
	// Description is a brief description of the step.
	Description = "Checks the authentication endpoint by sending a POST request " +
		"to it and verifying that the response is successful."
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

// Step represents the HTTP authentication step in the autograder process.
type Step struct {
	// cli is the HTTP client used to send requests.
	cli httpClient
	// user is the user to authenticate.
	user user.User
	// baseURL is the base URL of the application.
	baseURL string
	// password is the password of the user to authenticate.
	password string
}

// NewStep creates a new Step instance.
func NewStep(
	cli httpClient,
	baseURL string,
	usr user.User,
	password string,
) *Step {
	return &Step{
		cli:      cli,
		baseURL:  baseURL,
		user:     usr,
		password: password,
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
