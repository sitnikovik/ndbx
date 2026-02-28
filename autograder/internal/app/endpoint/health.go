package endpoint

// Health returns the URL for the health endpoint of the autograder application.
func (e Endpoint) Health() string {
	return e.baseURL + "/health"
}
