package endpoint

// EventLike returns the URL for the like endpoint of the specific event.
func (e Endpoint) EventLike(id string) string {
	return e.Event(id) + "/like"
}
