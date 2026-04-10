package endpoint

import (
	"context"
	"io"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

const (
	// Name is the name of the step.
	Name = "Sign-Up Step"
	// Description is a brief description of the step.
	Description = "Checks the sign-up endpoint by sending a POST request " +
		"to it and verifying that the response is successful" +
		" and finds out the user's id to the step variables"
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

// mongoClient defines the interface for interacting with MongoDB.
type mongoClient interface {
	// AllBy retrieves all documents from the specified collection that match the given key-value pairs.
	AllBy(
		ctx context.Context,
		collection string,
		by doc.KVs,
	) (doc.Documents, error)
}

// Step represents the HTTP sign-up step in the autograder process.
type Step struct {
	// cli is the HTTP client used to send requests.
	cli httpClient
	// mongo is the MongoDB client used to interact with the database.
	mongo mongoClient
	// user is the user to sign up.
	user user.User
	// baseURL is the base URL of the application.
	baseURL string
	// pwd is the password of the user.
	pwd string
}

// NewStep creates a new Step instance.
func NewStep(
	httpcli httpClient,
	mongo mongoClient,
	baseURL string,
	usr user.User,
	pwd string,
) *Step {
	return &Step{
		cli:     httpcli,
		mongo:   mongo,
		baseURL: baseURL,
		user:    usr,
		pwd:     pwd,
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
