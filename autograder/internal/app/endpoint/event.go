package endpoint

// Event returns the URL for the event endpoint of the autograder application.
func (e Endpoint) Event() string {
	return e.baseURL + "/event"
}
