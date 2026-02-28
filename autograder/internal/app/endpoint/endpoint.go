package endpoint

// Endpoint represents the base URL for the autograder application endpoints.
type Endpoint struct {
	// baseURL is the base URL for the autograder application endpoints.
	baseURL string
}

// NewEndpoint creates a new Endpoint instance with the given base URL.
func NewEndpoint(baseURL string) Endpoint {
	return Endpoint{
		baseURL: baseURL,
	}
}
