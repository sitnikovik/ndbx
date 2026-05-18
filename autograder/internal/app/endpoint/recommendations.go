package endpoint

// Recommendations returns the URL for the user's recommendations endpoint
// of the autograder application.
func (e Endpoint) Recommendations() string {
	return e.baseURL + "/recommendations"
}
