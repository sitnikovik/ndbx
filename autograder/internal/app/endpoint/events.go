package endpoint

// Events returns the URL for the events endpoint of the autograder application.
func (e Endpoint) Events() string {
	return e.baseURL + "/events"
}
