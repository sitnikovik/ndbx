package body

// Body represents the response body for the events endpoint.
type Body struct {
	// id is the unique identifier of the created event, extracted from the response body.
	id string
}

// NewBody creates a new Body instance with the given event ID.
func NewBody(id string) Body {
	return Body{
		id: id,
	}
}

// ID returns the event ID contained in the response body.
func (b Body) ID() string {
	return b.id
}
