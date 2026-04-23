package endpoint

// EventReviews returns the URL for the reviews endpoint of the specific event.
func (e Endpoint) EventReviews(id string) string {
	return e.Event(id) + "/reviews"
}
