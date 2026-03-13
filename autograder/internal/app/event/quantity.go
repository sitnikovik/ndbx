package event

// Quantity represents the quantity of attendees for an event.
type Quantity struct {
	// min is the minimum quantity of attendees for the event.
	min uint16
	// max is the maximum quantity of attendees for the event.
	max uint16
}

// NewQuantity creates a new Quantity instance.
func NewQuantity(minqty, maxqty uint16) Quantity {
	return Quantity{
		min: minqty,
		max: maxqty,
	}
}

// Max returns the maximum quantity of attendees for the event.
func (q Quantity) Max() int {
	return int(q.max)
}

// Min returns the minimum quantity of attendees for the event.
func (q Quantity) Min() int {
	return int(q.min)
}
