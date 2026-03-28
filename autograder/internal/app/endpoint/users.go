package endpoint

// Users returns the URL for the users endpoint of the autograder application.
func (e Endpoint) Users() string {
	return e.baseURL + "/users"
}
