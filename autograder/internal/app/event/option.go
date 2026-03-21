package event

// Option represents a functional option for configuring an Event.
type Option func(*Event)

// WithQuantity sets the quantity of attendees for the event in the Event.
func WithQuantity(q Quantity) Option {
	return func(b *Event) {
		b.qty = q
	}
}

// WithCosts sets the cost information related to the event.
func WithCosts(costs Costs) Option {
	return func(e *Event) {
		e.costs = costs
	}
}
