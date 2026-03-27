package endpoint

// User returns the URL for the specific user endpoint of the autograder application.
func (e Endpoint) User(id string) string {
	return e.baseURL + "/users/" + id
}
