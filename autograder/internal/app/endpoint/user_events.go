package endpoint

// UserEvents returns the URL for user's events endpoint.
func (e Endpoint) UserEvents(id string) string {
	return e.User(id) + "/events"
}
