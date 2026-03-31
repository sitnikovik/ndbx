package endpoint

// EventDislike returns the URL to dislike the specific event.
func (e Endpoint) EventDislike(id string) string {
	return e.Event(id) + "/dislike"
}
