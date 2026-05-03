package endpoint

// EventReview returns the URL for the review endpoint of the specific event.
func (e Endpoint) EventReview(eventID, reviewID string) string {
	return e.EventReviews(eventID) + "/" + reviewID
}
