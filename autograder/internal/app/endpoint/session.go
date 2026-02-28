package endpoint

// Session returns the URL for the session endpoint of the autograder application.
func (e Endpoint) Session() string {
	return e.baseURL + "/session"
}
