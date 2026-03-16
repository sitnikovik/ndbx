package endpoint

// SignUp returns the URL for the sign-up endpoint of the autograder application.
func (e Endpoint) SignUp() string {
	return e.baseURL + "/users"
}
