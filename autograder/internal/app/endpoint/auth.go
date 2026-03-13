package endpoint

// Auth returns the URL for the authentication endpoint of the autograder application.
func (e Endpoint) Auth() string {
	return e.baseURL + "/auth/login"
}
