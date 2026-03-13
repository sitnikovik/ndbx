package endpoint

// Logout returns the URL for the logout endpoint of the autograder application.
func (e Endpoint) Logout() string {
	return e.baseURL + "/auth/logout"
}
