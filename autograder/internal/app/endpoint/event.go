package endpoint

// Event returns the URL for the specific event endpoint of the autograder application.
func (e Endpoint) Event(id string) string {
	return e.baseURL + "/events/" + id
}
